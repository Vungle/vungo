package openrtbutil_test

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbutil"
)

// setupTestServer method starts up a test server and returns a dummy bid response. Caller may
// provide a middleware HandlerFunc to perform additional changes to the response before the bid
// response get written, but should not write to the response body.
// Caller is also responsible for closing the server when done.
func setupTestServer(t *testing.T, fn http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if h := req.Header.Get(openrtb.HeaderOpenRTBVersion); h != openrtb.OpenRTBVersion {
			t.Errorf("HTTP header for OpenRTB bid requests should contain `%s: %s` header instead of %s.", openrtb.HeaderOpenRTBVersion, openrtb.OpenRTBVersion, h)
		}

		resp.Header().Set("Content-Type", "application/json")
		f := "testdata/simple_bidresponse.json"
		data, err := ioutil.ReadFile(f)
		t.Log("Test bid response:\n" + string(data))
		if err != nil {
			t.Fatal("Failed to read test data from file: ", f)
		}

		if fn != nil {
			fn(resp, req)
		}

		resp.Write(data)
	}))
}

func TestClientDo(t *testing.T) {
	ts := setupTestServer(t, nil)
	defer ts.Close()

	resp, err := openrtbutil.MakeSimpleRequest(t, ts.URL)
	if err != nil {
		t.Fatal("Bid request should return a bid response instead of an error: ", err)
	}

	if resp == nil {
		t.Fatal("Bid response should be non-nil when there is no error.")
	}

	if resp.HTTP() == nil {
		t.Error("HTTP response should always be non-nil.")
	}

	if resp.OpenRTB() == nil {
		t.Error("OpenRTB bid response should always be non-nil.")
	}
}

func TestClientDoShouldNotOverrideExistingOpenRtbHeaders(t *testing.T) {
	// Setup expected value and done waiting.
	done := make(chan struct{})
	version := "2.0"

	// Setup a test server that would always return a no-bid.
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(http.StatusNoContent)
		if v := req.Header.Get(openrtb.HeaderOpenRTBVersion); v != version {
			t.Errorf("Expected request header to have override X-OpenRTB-Header to be %s rather than %s.", version, v)
		}
		close(done)
	}))
	defer ts.Close()

	// Make a custom request.
	ctx := context.Background()
	br := &openrtb.BidRequest{
		ID: "br-with-custom-header",
	}

	req, err := openrtbutil.NewRequest(ctx, br, ts.URL, nil)
	if err != nil {
		t.Fatal("Cannot create a bid request: ", err)
	}

	// Set custom OpenRTB version header.
	req.HTTP().Header.Set(openrtb.HeaderOpenRTBVersion, version)

	// Make the request to test server.
	c := openrtbutil.NewClient(nil)
	switch _, err := c.Do(req); err.(type) {
	case openrtbutil.NoBidError:
	default:
		t.Errorf("Expected error to be a NoBidError instead of %v.", err)
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Error("Test server took too long to respond, was there really a request?")
	}
}

func TestClientDoShouldContainExtraHeaders(t *testing.T) {
	header := "X-Custom-Test-Header"
	value := "我家住在黄土高坡~"
	ts := setupTestServer(t, func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set(header, value)
	})
	defer ts.Close()

	resp, err := openrtbutil.MakeSimpleRequest(t, ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if v := resp.HTTP().Header.Get(header); v != value {
		t.Errorf("Custom header %s should have value of `%s` instead of\n`%s`.", header, value, v)
	}
}

func TestClientDoShouldRespondNoBid(t *testing.T) {
	tmax := 100 * time.Millisecond
	tests := []struct {
		latency time.Duration
		header  string
		status  int
		payload []byte
	}{
		// When response takes too long.
		{120 * time.Millisecond, "application/json", http.StatusOK, []byte(`{"id":"abc","seatbid":[{"bid":[{"id":"1"}]}]}`)},

		// When response did not set proper Content-Type header.
		{0, "", http.StatusOK, []byte(`{"id":"abc","seatbid":[{"bid":[{}]}]}`)},

		// When response set the wrong Content-Type header.
		{0, "text/json", http.StatusOK, []byte(`{"id":"abc","seatbid":[{"bid":[{}]}]}`)},

		// When response set HTTP status to be 204 No Content.
		{0, "application/json", http.StatusNoContent, nil},

		// When response set HTTP status to be anything other than 200 status code.
		{0, "application/json", http.StatusBadRequest, []byte(`{"id":"abc","seatbid":[{"bid":[{}]}]}`)},
		{0, "application/json", http.StatusUnauthorized, []byte(`{"id":"abc","seatbid":[{"bid":[{}]}]}`)},
		{0, "application/json", http.StatusHTTPVersionNotSupported, []byte(`{"id":"abc","seatbid":[{"bid":[{}]}]}`)},

		// When response sent empty HTTP JSON body.
		{0, "application/json", http.StatusOK, []byte(`{}`)},

		// When response sent malformed JSON body:
		// 	non-JSON
		{0, "application/json", http.StatusOK, []byte(`this is not even JSON`)},
		{0, "application/json", http.StatusOK, []byte(`{id:"123"}`)},
		// 	id should be an string.
		{0, "application/json", http.StatusOK, []byte(`{"id":1,"seatbid":[]}`)},
		// 	seatbid should be an object array.
		{0, "application/json", http.StatusOK, []byte(`{"id":"abc","seatbid":{}}`)},
		// 	bid should be an object array.
		{0, "application/json", http.StatusOK, []byte(`{"id":1,"seatbid":["bid":{}]}`)},

		// When response sent JSON body without seatbid.
		{0, "application/json", http.StatusOK, []byte(`{"id":"abc"}`)},

		// When response sent JSON body with empty seatbid.
		{0, "application/json", http.StatusOK, []byte(`{"id":"abc","seatbid":[]}`)},
	}

	var middleware http.HandlerFunc
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		middleware(resp, req)
	}))
	defer ts.Close()

	for i, test := range tests {
		t.Logf("Testing %d...", i)

		middleware = func(resp http.ResponseWriter, req *http.Request) {
			if test.latency > 0 {
				time.Sleep(test.latency)
			}

			if len(test.header) > 0 {
				resp.Header().Set("Content-Type", test.header)
			}

			resp.WriteHeader(test.status)

			if len(test.payload) > 0 {
				resp.Write(test.payload)
			}
		}

		ctx, _ := context.WithTimeout(context.Background(), tmax)
		br := &openrtb.BidRequest{
			ID: "br-for-no-bid",
		}

		req, err := openrtbutil.NewRequest(ctx, br, ts.URL, nil)
		if err != nil {
			t.Fatal("Cannot create a bid request: ", err)
		}

		// Make the request to test server.
		c := openrtbutil.NewClient(nil)
		switch resp, err := c.Do(req); e := err.(type) {
		case net.Error:
			t.Log("No-bid from network error: ", e)
			if !e.Timeout() {
				t.Error("Returned network error should be a timeout.", err)
			}
		case openrtbutil.NoBidError:
			t.Log("No-bid from no-bid error: ", e)
			if resp != nil {
				t.Errorf("Response should always be nil when NoBidError was returned: %+v", resp)
			}
		default:
			t.Errorf("Expected error to be a NoBidError instead of %v.", err)
		}
	}
}
