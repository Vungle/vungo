package validator

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
)

type Vast3validator struct {
	version     vastbasic.Version
	validatorV2 *Vast2validator
}

// ValidateVersion method validates the Version according to the VAST.
func (vc Vast3validator) ValidateVersion(v vastbasic.Version) error {
	if v != vastbasic.Version3 {
		return ValidationError{Errs: []error{vastbasic.ErrUnsupportedVersion}}
	}

	return nil
}

// ValidateWrapper method validates the Wrapper according to the VAST.
// AdSystem, VastAdTagURI, and Impressions are required.
// Creatives are optional, if it exists, we'll also validate it.
func (vc *Vast3validator) ValidateWrapper(w *entity.Wrapper) error {
	errors := make([]error, 0)
	if len(w.VastAdTagURI) == 0 {
		errors = append(errors, vastbasic.ErrWrapperMissVastAdTagURI)
	}

	if len(w.Impressions) == 0 {
		errors = append(errors, vastbasic.ErrWrapperMissImpressions)
	}

	for _, c := range w.Creatives {
		if err := vc.validatorV2.ValidateCreative(c); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// ValidateVast method validates the Vast element according to the VAST.
// Version, and Ads are required.
func (vc *Vast3validator) ValidateVast(v *entity.Vast) error {
	errors := make([]error, 0)
	if err := vc.ValidateVersion(v.Version); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	if len(v.Ads) == 0 {
		errors = append(errors, vastbasic.ErrVastMissAd)
	}

	for _, ad := range v.Ads {
		if err := vc.validatorV2.ValidateAd(ad); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// more events

// ad pods

// adparameter isXmlEncoded

// icon

// mediafile
//Codec      string `xml:"codec,attr,omitempty"`      // VAST3.0.
//MinBitrate *int   `xml:"minBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.
//MaxBitrate *int   `xml:"maxBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.

// tracking offset

// ad Sequence

// Companion
//AssetWidth    int                    `xml:"assetWidth,attr"`                                // VAST3.0.
//AssetHeight   int                    `xml:"assetHeight,attr"`                               // VAST3.0.
//AdSlotID      string                 `xml:"adSlotId,attr,omitempty"`                        // VAST3.0.
//ClickTracking vastbasic.TrimmedData  `xml:"CompanionClickTracking,omitempty"`               // VAST3.0.
//Extensions    []*vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.

// CompanionAd Required

// creative APIFramework

// inline
//Advertiser string             `xml:"Advertiser,omitempty"` // VAST3.0.
//Pricing    *vastbasic.Pricing `xml:"Pricing,omitempty"`    // VAST3.0.

// Linear
//Extensions []*vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
//Icons      []*vastbasic.Icon      `xml:"Icons>Icon"`                                     // VAST3.0.

// nolinear
//ClickTracking []string              `xml:"NonLinearClickTracking,omitempty"`               // VAST3.0.
//Extensions    []vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
