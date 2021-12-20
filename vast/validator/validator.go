package validator

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/entity"
)

// Validator is used as the abstract function of vast struct validate tool
type Validator interface {
	ValidateVersion(v vastbasic.Version) error
	ValidateAdSystem(as *vastbasic.AdSystem) error
	ValidateDelivery(delivery vastbasic.Delivery) error
	ValidateDuration(d vastbasic.Duration) error
	ValidateEvent(e vastbasic.Event) error
	ValidateIcon(icon *vastbasic.Icon) error
	ValidateImpression(impression *vastbasic.Impression) error
	ValidateMediaFile(mediaFile *vastbasic.MediaFile) error
	ValidateMode(mode vastbasic.Mode) error
	ValidateOffset(o *vastbasic.Offset) error
	ValidatePricing(pricing *vastbasic.Pricing) error
	ValidatePricingModel(module vastbasic.PricingModel) error
	ValidateStaticResource(sr *vastbasic.StaticResource) error
	ValidateTracking(t *vastbasic.Tracking) error
	ValidateVideoClick(videoClick *vastbasic.VideoClick) error
	ValidateVideoClicks(videoClicks *vastbasic.VideoClicks) error
	ValidateAd(ad *entity.Ad) error
	ValidateCompanion(companion *entity.Companion) error
	ValidateCompanionAds(companionAds *entity.CompanionAds) error
	ValidateCreative(creative *entity.Creative) error
	ValidateInLine(inline *entity.InLine) error
	ValidateLinear(linear *entity.Linear) error
	ValidateNonLinear(nonLinear *entity.NonLinear) error
	ValidateNonLinearAds(nonLinearAds *entity.NonLinearAds) error
	ValidateWrapper(w *entity.Wrapper) error
	ValidateVast(v *entity.Vast) error
}
