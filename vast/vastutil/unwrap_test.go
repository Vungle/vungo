package vastutil_test

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vastutil"
)

func TestUnwrap(t *testing.T) {
	tc := newTestUnwrapClient()
	vastutil.UpdateDefaultUnwrapClient(tc.Client)
	defer vastutil.RevertDefaultUnwrapClient()

	// The tests performs unwrapping of VAST XML starting from the entry point and is expected to
	// follow the trace until an error occurs or succeeds.
	tests := []struct {
		desc          string   // description for the test case
		entry         string   // entry point for the VAST unwrapping test.
		trace         []string // list of steps traces through.
		expectedError error    // expected error in case error occurs.
	}{
		{"vast element with error xml format", "malformed1", []string{}, &xml.SyntaxError{Line: 4, Msg: "unexpected EOF"}},
		{"valid inline vast element", "inline1", []string{}, nil},
		{"valid wrapper vast element", "wrapper3ads", []string{"inline1"}, nil},
		{"invalid wrapper without TagURI", "wrapper4nouri", []string{}, vastutil.ErrWrapperMissingAdTagURI},
		{"invalid xml format in wrapper response", "wrapper1", []string{"malformed1"}, &xml.SyntaxError{Line: 4, Msg: "unexpected EOF"}},
		{"valid inline in wrapper response", "wrapper1", []string{"inline1"}, nil},
		{"valid inline in the second wrapper response", "wrapper1", []string{"wrapper2", "inline1"}, nil},
		{"valid triple wrapper response with one empty vast element", "unwrapper_multiads", []string{"unwrapper_multiads_1", "unwrapper_multiads_2", "unwrapper_multiads_3"}, nil},
	}

	ctx := context.Background()

	for i, test := range tests {
		t.Logf("Testing %d starting from %s...", i, test.entry)
		tc.Init(test.trace)
		data, err := ioutil.ReadFile(testFilePath(test.entry))

		if err != nil {
			t.Fatal(err)
		} else if err = tc.AddToServed(testFilePath(test.entry)); err != nil {
			t.Fatal(err)
		}

		vasts, err := vastutil.Unwrap(ctx, data, "fake-ua", "1.2.3.4")

		if err != nil && !reflect.DeepEqual(err, test.expectedError) {
			t.Log(reflect.TypeOf(err))
			t.Errorf("Failed for %s. Expecting error %v instead of %v.", test.desc, test.expectedError, err)
		} else if err == nil {
			if len(tc.served) != len(test.trace)+1 {
				t.Errorf("Failed for %s. Unwrapping VAST should hop %d times instead of %d times.", test.desc, len(test.trace)+1, len(tc.served))
			} else if len(vasts) != len(tc.served) {
				t.Errorf("Failed for %s. Unwrapped VAST depth %d was different from served vast, %d.", test.desc, len(vasts), len(tc.served))
			} else if !reflect.DeepEqual(vasts, tc.served) {
				t.Errorf("Failed for %s. Unwrapped VAST does not match with served VAST. unwrapped: %v, served: %v", test.desc, vasts, tc.served)
			}
		}
	}
}

func TestUnwrapShouldRespectContext(t *testing.T) {
	// Given a context that is already finished.
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// When unwraps a wrapper VAST XML.
	data, err := ioutil.ReadFile(testFilePath("wrapper1"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = vastutil.Unwrap(ctx, data, "fake-ua", "1.2.3.4")

	// Expect error to occur.
	if !strings.Contains(err.Error(), context.Canceled.Error()) {
		t.Errorf("Expected %v, got %v.", context.Canceled, err)
	}
}

// testFilePath method returns the test VAST XML file path based on the test name.
func testFilePath(name string) string {
	return fmt.Sprintf("testdata/unwrap/%s.xml", name)
}

// testUnwrapClient type is a http.Client wrapper that mocks the actual HTTP client used during the
// unwrapping process. The testUnwrapClient will swallow the actual HTTP request and respond with
// the test XML file.
type testUnwrapClient struct {
	*http.Client

	served []*vastelement.Vast

	at    int
	trace []string
}

// Init method initializes the client with a slice of XML files to respond with. These XML responses
// are sequential, so the order matters.
func (c *testUnwrapClient) Init(trace []string) {
	c.served = make([]*vastelement.Vast, 0)
	c.at = 0
	c.trace = trace
}

// RoundTrip implements the http.RoundTripper interface to mock out the actual HTTP request.
func (c *testUnwrapClient) RoundTrip(_ *http.Request) (*http.Response, error) {
	resp := &http.Response{}
	f := testFilePath(c.trace[c.at])
	body, err := os.Open(f)

	c.at++

	if err != nil {
		return nil, err
	} else if err := c.AddToServed(f); err != nil {
		return nil, err
	}

	resp.Body = body
	return resp, nil
}

// AddToServed method adds the actual vast.Vast object to a list for testing purposes.
func (c *testUnwrapClient) AddToServed(f string) error {
	var v vastelement.Vast
	if data, err := ioutil.ReadFile(f); err != nil {
		return err
	} else if err := xml.Unmarshal(data, &v); err != nil {
		// Swallow this error as it should be surfaced by the actual test.
		return nil
	} else {
		c.served = append(c.served, &v)
		return nil
	}
}

// newTestUnwrapClient method instantiates a new testing client.
func newTestUnwrapClient() *testUnwrapClient {
	c := &testUnwrapClient{Client: &http.Client{}}
	c.Transport = c
	return c
}
