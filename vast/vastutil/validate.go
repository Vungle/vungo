package vastutil

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/Vungle/vungo/vast/vastelement"
)

// Validate method validate VAST data as an external api
func Validate(vasts []*vastelement.Vast, version vastelement.Version) error {
	var impressionsErrCount int
	var hasLinear bool
	for _, v := range vasts {
		if err := v.Validate(version); err != nil {
			invalidVast, marshalErr := xml.Marshal(v)
			if marshalErr != nil {
				return fmt.Errorf("failed to marshal VAST for logging: %v", marshalErr)
			}
			if strings.Contains(err.Error(), vastelement.ErrWrapperMissImpressions.Error()) &&
				impressionsErrCount < len(vasts)-1 {
				impressionsErrCount++
			} else {
				return fmt.Errorf("failed to validate VAST: %v, VAST: %s", err.Error(), invalidVast)
			}
		}
		linear := v.FindFirstInlineLinearCreative()
		if linear == nil || len(linear.MediaFiles) == 0 {
			continue
		}
		hasLinear = true
	}
	if !hasLinear {
		return fmt.Errorf("vasts should have at least one linear")
	}

	return nil
}
