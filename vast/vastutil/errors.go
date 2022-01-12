package vastutil

import "errors"

// Please keep the errors alphabetized.

// ErrUnwrapWithMultipleAds error represents the error during unwrapping when the VAST contains
// multiple Ad elements.
var ErrUnwrapWithMultipleAds = errors.New("unsupported unwrapping with multiple ads")

// ErrWrapperMissingAdTagURI error represents when the wrapper VAST xml is missing required
// VASTAdTagURI element.
var ErrWrapperMissingAdTagURI = errors.New("invalid VAST content, missing ad tag URI")
