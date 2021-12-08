package vastutil

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Vungle/vungo/vast/vast2"
)

var defaultUnwrapClient = &http.Client{
	Transport: &http.Transport{
		// Initialize TLSNextProto to disable HTTP/2 support.
		// For more details please refer to https://github.com/golang/go/issues/32388
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	},
}

// Unwrap method recursively unmarshals the VAST XML data from the input and unwraps additional VAST
// XML content when the current VAST is a wrapper. This method returns a slice of all the VAST XML
// unwrapped during the process and exactly one Inline VAST; or error on any of the following
// conditions:
// 	1) Unwrapping context, ctx, is done;
// 	2) Invalid XML format;
// 	3) Invalid VAST content, (currently, there should be exactly one <Ad> within <VAST>.
func Unwrap(ctx context.Context, data []byte, userAgent, ip string) ([]*vast2.Vast, error) {
	var v vast2.Vast
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	return unwrap(ctx, &v, nil, userAgent, ip)
}

// unwrap method is a helper function that invokes performs the actual HTTP request to get
// additional VAST XML and updates the unwrappedList. The unwrap method is invoked recursively until
// the first Inline VAST is reached or until an error occurs.
func unwrap(ctx context.Context, v *vast2.Vast, unwrappedList []*vast2.Vast, ua, ip string) ([]*vast2.Vast, error) {
	if len(v.Ads) != 1 {
		return nil, ErrUnwrapWithMultipleAds
	} else if v.Ads[0].Wrapper == nil {
		return append(unwrappedList, v), nil
	}

	w := v.Ads[0].Wrapper
	if len(w.VastAdTagURI) == 0 {
		return nil, ErrWrapperMissingAdTagURI
	}

	var innerVast vast2.Vast

	req, err := http.NewRequest("GET", w.VastAdTagURI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	req.Header.Set("X-Forwarded-For", ip)
	req = req.WithContext(ctx)
	resp, err := defaultUnwrapClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	if err = xml.NewDecoder(resp.Body).Decode(&innerVast); err != nil {
		return nil, err
	}
	// TODO(@garukun): Given a set of super fast VAST hosts and a starting wrapper VAST XML that
	// wraps the VAST infinitely, this implementation could cause disastrous stack overflow that
	// even the ctx.Done() cannot enforce exit. We will need to update the logic here eventually
	// to be more resource aware yet robust. *wink* *wink*, consider channels and goroutines.
	return unwrap(ctx, &innerVast, append(unwrappedList, v), ua, ip)
}
