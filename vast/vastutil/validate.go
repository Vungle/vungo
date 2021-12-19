package vastutil

import (
	"encoding/xml"
	"fmt"
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
	"github.com/Vungle/vungo/vast/validator"
	"strings"
)

func Validate(vasts []*entity.Vast, version vastbasic.Version) error {
	validator := validatorOf(version)
	if validator == nil {
		return fmt.Errorf("[vastutil Validate] fail to get validator (version: %s)", version)
	}
	var impressionsErrCount int
	var hasLinear bool
	for _, v := range vasts {
		if err := validator.ValidateVast(v); err != nil {
			invalidVast, marshalErr := xml.Marshal(v)
			if marshalErr != nil {
				return fmt.Errorf("failed to marshal VAST for logging: %v", marshalErr)
			}
			if strings.Contains(err.Error(), vastbasic.ErrWrapperMissImpressions.Error()) &&
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

func validatorOf(version vastbasic.Version) validator.Validator {
	switch version {
	case vastbasic.Version2:
		return &validator.Vast2validator{}
	case vastbasic.Version3:
		return &validator.Vast3validator{}
	default:
		return nil
	}
}
