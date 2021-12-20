package validator

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/defaults"
	"github.com/Vungle/vungo/vast/entity"
)

// Vast3validator is a validate tool for vast 3.0.
// new feature of Vast 3.0 are listed below:
//		More events: EventExitFullscreen EventSkip EventProgress
//		- Ad pods and
//			sequence			int 			new attribute
//		- Adparameter
//			isXmlEncoded 		bool			new attribute
//		- Icon
//			Program      		string   		new attribute
//			Width        		int  			new attribute
//			Height       		int     		new attribute
//			XPosition   		string   		new attribute
//			YPosition    		string 			new attribute
//			Duration     		Duration		new attribute
//			APIFramework 		string  		new element
//			Offset       		Offset   		new attribute
//			ClickThrough   		string        	new element
//			ClickTrackings 		[]string        new element
//			StaticResource 		*StaticResource new element
//			IFrameResource 		string         	new element
//			HTMLResource   		*HTMLResource   new element
//		- Mediafile
//			Codec				string			new attribute
//			MinBitrate			*int			new attribute
//			MaxBitrate			*int			new attribute
//		- Tracking
//			offset				*Offset			new attribute
//		- Companion
//			AssetWidth			int				new attribute
//			AssetHeight			int				new attribute
//			AdSlotID			string			new attribute
//			ClickTracking		TrimmedData		new element
//			Extensions    		*Extension 		new element
//		- CompanionAd
//			Required			Mode			new attribute
//		- Creative
//			APIFramework		string			new attribute
//		- Inline
//			Advertiser	 		string      	new element
//			Pricing    			*Pricing 		new element
//		- Linear
//			Extensions 			[]*Extension 	new element
//			Icons      			[]*Icon			new element
//		- Nolinear
//			ClickTracking		[]string		new element
//			Extensions			[]Extension		new element
type Vast3validator struct {
}

// ValidateVersion method validates the Version according to the VAST.
func (vc Vast3validator) ValidateVersion(v vastbasic.Version) error {
	if v != vastbasic.Version3 {
		return ValidationError{Errs: []error{vastbasic.ErrUnsupportedVersion}}
	}

	return nil
}

// ValidateAdSystem method validate AdSystem vast element
func (vc *Vast3validator) ValidateAdSystem(as *vastbasic.AdSystem) error {
	if len(as.System) == 0 {
		return ValidationError{Errs: []error{vastbasic.ErrAdSystemMissSystem}}
	}
	return nil
}

// ValidateDelivery method validate Delivery vast element
func (vc *Vast3validator) ValidateDelivery(delivery vastbasic.Delivery) error {
	if delivery != vastbasic.DeliveryProgressive && delivery != vastbasic.DeliveryStreaming {
		return ValidationError{Errs: []error{vastbasic.ErrUnsupportedDeliveryType}}
	}
	return nil
}

// ValidateDuration method validate Duration vast element
func (vc *Vast3validator) ValidateDuration(d vastbasic.Duration) error {
	if d < 0 {
		return ValidationError{Errs: []error{vastbasic.ErrDurationNegative}}
	}
	if d == 0 {
		return ValidationError{Errs: []error{vastbasic.ErrDurationEqualZero}}
	}
	return nil
}

// ValidateEvent method validate Event vast element
func (vc *Vast3validator) ValidateEvent(e vastbasic.Event) error {
	switch e {
	case vastbasic.EventCreativeView:
	case vastbasic.EventStart:
	case vastbasic.EventFirstQuartile:
	case vastbasic.EventMidpoint:
	case vastbasic.EventThirdQuartile:
	case vastbasic.EventComplete:
	case vastbasic.EventMute:
	case vastbasic.EventUnmute:
	case vastbasic.EventPause:
	case vastbasic.EventRewind:
	case vastbasic.EventResume:
	case vastbasic.EventFullscreen:
	case vastbasic.EventExitFullscreen:
	case vastbasic.EventExpand:
	case vastbasic.EventCollapse:
	case vastbasic.EventAcceptInvitation:
	case vastbasic.EventClose:
	case vastbasic.EventSkip:
	case vastbasic.EventProgress:
	default:
		// Validate function returns error of event.
		// We do not do the strict validation. If the event type is not defined, just skipped it rather than return validate failure.
		// Please refer to https://vungle.atlassian.net/browse/PBJ-685 for more details.
		return nil
	}
	return nil
}

// ValidateIcon method validate Icon vast element
func (vc *Vast3validator) ValidateIcon(icon *vastbasic.Icon) error {
	errors := make([]error, 0)
	if len(icon.Program) == 0 {
		errors = append(errors, vastbasic.ErrIconMissProgram)
	}

	if len(icon.XPosition) == 0 || len(icon.YPosition) == 0 {
		errors = append(errors, vastbasic.ErrIconMissPosition)
	}

	if icon.StaticResource != nil {
		if len(icon.IFrameResource) != 0 || icon.HTMLResource != nil {
			errors = append(errors, vastbasic.ErrIconResourcesFormat)
		}
	} else if icon.HTMLResource != nil {
		if len(icon.IFrameResource) != 0 {
			errors = append(errors, vastbasic.ErrIconResourcesFormat)
		}
	} else if len(icon.IFrameResource) == 0 {
		errors = append(errors, vastbasic.ErrIconResourcesFormat)
	}
	if icon.StaticResource != nil {
		if err := vc.ValidateStaticResource(icon.StaticResource); err != nil {
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

// ValidateImpression method validate Impression vast element
func (vc *Vast3validator) ValidateImpression(impression *vastbasic.Impression) error {
	if len(impression.URI) == 0 {
		return ValidationError{Errs: []error{vastbasic.ErrImpressionMissURI}}
	}
	return nil
}

// ValidateMediaFile method validate the MediaFile element according to the VAST.
// Delivery and MimeType, Width, and Height are required.
// Since the width, and height might be zero, we'll only make sure they are not less than zero.
func (vc *Vast3validator) ValidateMediaFile(mediaFile *vastbasic.MediaFile) error {
	errors := make([]error, 0)

	mimeTypeIsSupported := false
	for _, mimeType := range defaults.SupportedMineTypes {
		if mimeType == mediaFile.MimeType {
			mimeTypeIsSupported = true
			break
		}
	}
	if !mimeTypeIsSupported {
		errors = append(errors, vastbasic.ErrMediaFileUnsupportedMimeType)
		return ValidationError{Errs: errors}
	}

	if len(mediaFile.Delivery) == 0 {
		errors = append(errors, vastbasic.ErrMediaFileMissDelivery)
	}

	if err := vc.ValidateDelivery(mediaFile.Delivery); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	if mediaFile.Width < 0 || mediaFile.Height < 0 {
		errors = append(errors, vastbasic.ErrMediaFileSize)
	}

	if len(mediaFile.URI) == 0 {
		errors = append(errors, vastbasic.ErrMediaFileMissURI)
	}

	if mediaFile.Width > defaults.MaxVideoWidth {
		errors = append(errors, vastbasic.ErrMediaFileWidthTooHigh)
	}

	if mediaFile.Width < defaults.MinVideoWidth {
		errors = append(errors, vastbasic.ErrMediaFileWidthTooLow)
	}

	if mediaFile.Height > defaults.MaxVideoHeight {
		errors = append(errors, vastbasic.ErrMediaFileHeightTooHigh)
	}

	if mediaFile.Height < defaults.MinVideoHeight {
		errors = append(errors, vastbasic.ErrMediaFileHeightTooLow)
	}

	if mediaFile.MinBitrate != nil && *(mediaFile.MinBitrate) < 0 {
		errors = append(errors, vastbasic.ErrMediaMinBitrateLessThanZero)
	}

	if mediaFile.MaxBitrate != nil && *(mediaFile.MaxBitrate) < 0 {
		errors = append(errors, vastbasic.ErrMediaMaxBitrateLessThanZero)
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// ValidateMode method validates Mode according to the VAST.
func (vc *Vast3validator) ValidateMode(mode vastbasic.Mode) error {
	if mode != vastbasic.ModeAll && mode != vastbasic.ModeAny && mode != vastbasic.ModeNone {
		return ValidationError{Errs: []error{vastbasic.ErrUnsupportedMode}}
	}

	return nil
}

// ValidateOffset method validate Offset.
func (vc *Vast3validator) ValidateOffset(o *vastbasic.Offset) error {
	errors := make([]error, 0)
	if o.Percent < 0 {
		errors = append(errors, vastbasic.ErrOffsetPercentNegative)
	}

	if err := vc.ValidateDuration(o.Duration); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// ValidatePricing methods validates the Pricing according to the VAST.
// Model, Price is required. Currency should be ISO-4217 3 letter format.
func (vc *Vast3validator) ValidatePricing(pricing *vastbasic.Pricing) error {
	return nil
}

// ValidatePricingModel method validates the PricingModel according to the VAST.
func (vc *Vast3validator) ValidatePricingModel(module vastbasic.PricingModel) error {

	switch module {
	case vastbasic.PricingModelCPM:
	case vastbasic.PricingModelCPC:
	case vastbasic.PricingModelCPE:
	case vastbasic.PricingModelCPV:
	default:
		return ValidationError{Errs: []error{vastbasic.ErrUnsupportedPriceModelType}}
	}
	return nil
}

// ValidateStaticResource method validates the StaticResource according to the VAST.
func (vc *Vast3validator) ValidateStaticResource(sr *vastbasic.StaticResource) error {
	return nil
}

// ValidateTracking methods validate the Tracking element according to the VAST.
func (vc *Vast3validator) ValidateTracking(t *vastbasic.Tracking) error {
	errors := make([]error, 0)

	if err := vc.ValidateEvent(t.Event); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// ValidateVideoClick method validates VideoClick according to the VAST.
func (vc *Vast3validator) ValidateVideoClick(videoClick *vastbasic.VideoClick) error {
	return nil
}

// ValidateVideoClicks method validates the VideoClicks according to the VAST.
// ClickThroughs element is required.
func (vc *Vast3validator) ValidateVideoClicks(videoClicks *vastbasic.VideoClicks) error {
	errors := make([]error, 0)

	// TODO(@DevKai): VAST3.0 require ClickThroughs not be empty.

	for _, click := range videoClicks.ClickThroughs {
		if err := vc.ValidateVideoClick(click); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, click := range videoClicks.ClickTrackings {
		if err := vc.ValidateVideoClick(click); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, click := range videoClicks.CustomClicks {
		if err := vc.ValidateVideoClick(click); err != nil {
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

// ValidateAd methods validate the Ad element according to the VAST.
// Each <Ad> contains either EXACTLY ONE <InLine> element or <Wrapper> element (but never both).
func (vc *Vast3validator) ValidateAd(ad *entity.Ad) error {
	errors := make([]error, 0)

	if ad.InLine != nil && ad.Wrapper != nil {
		errors = append(errors, vastbasic.ErrAdType)
		return ValidationError{Errs: errors}
	}

	if ad.InLine != nil {
		errors = append(errors, vc.ValidateAdInline(ad)...)
	} else if ad.Wrapper != nil {
		errors = append(errors, vc.ValidateAdWrapper(ad)...)
	} else {
		errors = append(errors, vastbasic.ErrAdType)
	}

	if ad.Sequence < 0 {
		errors = append(errors, vastbasic.ErrAdSequenceLessThanZero)
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// ValidateAdInline method validate Inline in an Ad vast element
func (vc *Vast3validator) ValidateAdInline(ad *entity.Ad) []error {
	errors := make([]error, 0)
	if err := vc.ValidateInLine(ad.InLine); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	return errors
}

// ValidateAdWrapper method validate Wrapper in an Ad vast element
func (vc *Vast3validator) ValidateAdWrapper(ad *entity.Ad) []error {
	errors := make([]error, 0)
	if err := vc.ValidateWrapper(ad.Wrapper); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	return errors
}

// ValidateCompanion methods validate the Companion element according to the VAST.
func (vc *Vast3validator) ValidateCompanion(companion *entity.Companion) error {
	errors := make([]error, 0)

	for _, tracking := range companion.Trackings {
		if err := vc.ValidateTracking(tracking); err != nil {
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

// ValidateCompanionAds methods validate the CompanionAds element according to the VAST.
func (vc *Vast3validator) ValidateCompanionAds(companionAds *entity.CompanionAds) error {
	errors := make([]error, 0)

	// No need to validate Required which is for VAST 3.0 only.

	for _, companion := range companionAds.Companions {
		if err := vc.ValidateCompanion(companion); err != nil {
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

// ValidateCreative methods validate the Creative element according to the VAST.
// Creative must contain EXACTLY ONE of Linear, CompanionAds, or NonLinearAds.
func (vc *Vast3validator) ValidateCreative(creative *entity.Creative) error {
	// Linear, CompanionAds, NonLinearAds are all optional?
	errors := make([]error, 0)
	if creative.Linear != nil {
		if creative.CompanionAds != nil || creative.NonLinearAds != nil {
			errors = append(errors, vastbasic.ErrCreativeType)
		}
		if err := vc.ValidateLinear(creative.Linear); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	} else if creative.CompanionAds != nil {
		if creative.NonLinearAds != nil {
			errors = append(errors, vastbasic.ErrCreativeType)
		}
		if err := vc.ValidateCompanionAds(creative.CompanionAds); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	} else if creative.NonLinearAds == nil {
		errors = append(errors, vastbasic.ErrCreativeType)
	} else if creative.NonLinearAds != nil {
		if err := vc.ValidateNonLinearAds(creative.NonLinearAds); err != nil {
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

// ValidateInLine methods validate the Inline element according to the VAST.
// AdSystem, AdTitle, Impression, Creatives are required.
func (vc *Vast3validator) ValidateInLine(inline *entity.InLine) error {
	errors := make([]error, 0)
	if inline.AdTitle == nil {
		errors = append(errors, vastbasic.ErrInlineMissAdTitle)
	}

	if len(inline.Creatives) == 0 {
		errors = append(errors, vastbasic.ErrInlineMissCreatives)
	}

	// No need to validate Pricing which is for VAST 3.0 only.

	if len(inline.Impressions) == 0 {
		errors = append(errors, vastbasic.ErrInlineMissImpressions)
	}

	for _, creative := range inline.Creatives {
		if err := vc.ValidateCreative(creative); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
		if creative.Linear != nil {
			if err := vc.ValidateDuration(creative.Linear.Duration); err != nil {
				return err
			}

			if creative.Linear.Duration > vastbasic.Duration(defaults.MaxVideoDuration) {
				errors = append(errors, vastbasic.ErrVideoDurationTooLong)
			}

			if creative.Linear.Duration < vastbasic.Duration(defaults.MinVideoDuration) {
				errors = append(errors, vastbasic.ErrVideoDurationTooShort)
			}

			if len(creative.Linear.MediaFiles) == 0 {
				errors = append(errors, vastbasic.ErrLinearMissMediaFiles)
			}

			var validMediaFiles []*vastbasic.MediaFile
			var err error

			var hasMimeTypeErr bool
			noneMimeTypeErrors := make([]error, 0)

			for _, mediaFile := range creative.Linear.MediaFiles {
				err = vc.ValidateMediaFile(mediaFile)
				if err == nil {
					validMediaFiles = append(validMediaFiles, mediaFile)
					//break
				} else {
					ve, ok := err.(ValidationError)
					if ok {
						// Merge all errors which are not mime type errors.
						if ve.Errs[0] != vastbasic.ErrMediaFileUnsupportedMimeType {
							noneMimeTypeErrors = append(noneMimeTypeErrors, ve.Errs...)
						} else {
							hasMimeTypeErr = true
						}
					}
				}
			}

			if len(validMediaFiles) > 0 {
				creative.Linear.MediaFiles = validMediaFiles
			} else {
				if len(noneMimeTypeErrors) > 0 {
					errors = append(errors, noneMimeTypeErrors...)
				}
				if hasMimeTypeErr {
					errors = append(errors, vastbasic.ErrMediaFileUnsupportedMimeType)
				}
			}

			if creative.Linear.Icons != nil && len(creative.Linear.Icons) > 0 {
				for _, icon := range creative.Linear.Icons {
					err = vc.ValidateIcon(icon)
					if err != nil {
						errors = append(errors, err)
					}
				}
			}
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// ValidateLinear methods validate the Linear element according to the VAST.
// Duration are required in Linear but not required in Wrapper.
// According to VAST2.0 protocol, MediaFile is not required.
// Icons are optional, but if contained, we'll also validate it.
func (vc Vast3validator) ValidateLinear(linear *entity.Linear) error {
	errors := make([]error, 0)

	if linear.VideoClicks != nil {
		if err := vc.ValidateVideoClicks(linear.VideoClicks); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if linear.SkipOffset != nil {
		if err := vc.ValidateOffset(linear.SkipOffset); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	// No need to validate Icon which is for VAST 3.0 only.

	for _, tracking := range linear.Trackings {
		if err := vc.ValidateTracking(tracking); err != nil {
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

// ValidateNonLinear a non linear struct
func (vc Vast3validator) ValidateNonLinear(nonLinear *entity.NonLinear) error {
	errors := make([]error, 0)
	// TODO(@DevKai): In VAST3.0, NonLinear resources should contain only one of StaticResource, IFrameResource, HTMLResource.
	if nonLinear.StaticResource != nil {
		if err := vc.ValidateStaticResource(nonLinear.StaticResource); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}
	if nonLinear.MinSuggestedDuration != nil {
		if err := vc.ValidateDuration(*nonLinear.MinSuggestedDuration); err != nil && err != vastbasic.ErrDurationEqualZero {
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

// ValidateNonLinearAds method validates the NonLinearAds according to the VAST.
// NonLinears is required.
func (vc *Vast3validator) ValidateNonLinearAds(nonLinearAds *entity.NonLinearAds) error {
	errors := make([]error, 0)
	if len(nonLinearAds.NonLinears) == 0 {
		errors = append(errors, vastbasic.ErrNonLinearAdsMissNonLinears)
	}

	for _, nonLinear := range nonLinearAds.NonLinears {
		if err := vc.ValidateNonLinear(nonLinear); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, tracking := range nonLinearAds.Trackings {
		if err := vc.ValidateTracking(tracking); err != nil {
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

// ValidateWrapper method validates the Wrapper according to the VAST.
// AdSystem, VastAdTagURI, and Impressions are required.
// Creatives are optional, if it exists, we'll also Validate it.
func (vc *Vast3validator) ValidateWrapper(w *entity.Wrapper) error {
	errors := make([]error, 0)
	if len(w.VastAdTagURI) == 0 {
		errors = append(errors, vastbasic.ErrWrapperMissVastAdTagURI)
	}

	if len(w.Impressions) == 0 {
		errors = append(errors, vastbasic.ErrWrapperMissImpressions)
	}

	for _, c := range w.Creatives {
		if err := vc.ValidateCreative(c); err != nil {
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
		if err := vc.ValidateAd(ad); err != nil {
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