package vastutil

import (
	"encoding/xml"
	"net/http"

	"github.com/Vungle/vungo/vast"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

var defaultUnwrapClient = http.DefaultClient

// Unwrap method recursively unmarshals the VAST XML data from the input and unwraps additional VAST
// XML content when the current VAST is a wrapper. This method returns a slice of all the VAST XML
// unwrapped during the process and exactly one Inline VAST; or error on any of the following
// conditions:
// 	1) Unwrapping context, ctx, is done;
// 	2) Invalid XML format;
// 	3) Invalid VAST content, (currently, there should be exactly one <Ad> within <VAST>.
func Unwrap(ctx context.Context, data []byte) ([]*vast.Vast, error) {
	var v vast.Vast
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	return unwrap(ctx, &v, nil)
}

// unwrap method is a helper function that invokes performs the actual HTTP request to get
// additional VAST XML and updates the unwrappedList. The unwrap method is invoked recursively until
// the first Inline VAST is reached or until an error occurs.
func unwrap(ctx context.Context, v *vast.Vast, unwrappedList []*vast.Vast) ([]*vast.Vast, error) {
	if len(v.Ads) != 1 {
		return nil, ErrUnwrapWithMultipleAds
	} else if v.Ads[0].Wrapper == nil {
		return append(unwrappedList, v), nil
	}

	w := v.Ads[0].Wrapper
	if len(w.VastAdTagUri) == 0 {
		return nil, ErrWrapperMissingAdTagUri
	}

	var innerVast vast.Vast
	resp, err := ctxhttp.Get(ctx, defaultUnwrapClient, w.VastAdTagUri)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		return nil, err
	} else if err = xml.NewDecoder(resp.Body).Decode(&innerVast); err != nil {
		return nil, err
	} else {
		// TODO(@garukun): Given a set of super fast VAST hosts and a starting wrapper VAST XML that
		// wraps the VAST infinitely, this implementation could cause disastrous stack overflow that
		// even the ctx.Done() cannot enforce exit. We will need to update the logic here eventually
		// to be more resource aware yet robust. *wink* *wink*, consider channels and goroutines.
		return unwrap(ctx, &innerVast, append(unwrappedList, v))
	}
}
