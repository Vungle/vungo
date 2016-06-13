package vastutil

import "errors"

// Please keep the errors alphabetized.

// ErrUnwrapWithMultipleAds error represents the error during unwrapping when the VAST contains
// multiple Ad elements.
var ErrUnwrapWithMultipleAds = errors.New("Unsupported unwrapping with multiple ads.")

// ErrWrapperMissingAdTagUri error represents when the wrapper VAST xml is missing required
// VASTAdTagURI element.
var ErrWrapperMissingAdTagUri = errors.New("Invalid VAST content, missing ad tag URI.")
