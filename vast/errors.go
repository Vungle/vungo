package vast

import "errors"

// Please keep the variables alphabetized.
var (
	ErrAdSystemMissSystem = errors.New("AdSystem misses system")

	ErrAdType = errors.New("Ad should only contain one of Inline and Wrapper")

	ErrCompanionAdsMissCompanions = errors.New("CompanionAds miss Companions")

	ErrCompanionAdsWrapperMissCompanions = errors.New("CompanionAdsWrapper misses Companions")

	ErrCreativeType = errors.New("Creative should contain one type of NonLinearAds, Linear, CompanionAds")

	ErrCreativeWrapperType = errors.New("CreativeWrapper should contain one type of NonLinearAds, Linear, CompanionAds")

	ErrDurationEqualZero = errors.New("Duration is less than zero")

	ErrDurationNegative = errors.New("Duration is negative")

	ErrIconMissPosition = errors.New("Icon misses position")

	ErrIconMissProgram = errors.New("Icon misses program")

	ErrIconResourcesFormat = errors.New("Icon resources should contain only one of StaticResource, IFrameResource, HtmlResource")

	ErrImpressionMissUri = errors.New("Impression misses Uri")

	ErrInlineMissAdSystem = errors.New("Inline misses AdSystem")

	ErrInlineMissAdTitle = errors.New("Inline misses AdTitle")

	ErrInlineMissCreatives = errors.New("Inline AD misses Creatives")

	ErrInlineMissImpressions = errors.New("Inline misses Impressions")

	ErrLinearMissMediaFiles = errors.New("MediaFiles are missed")

	ErrMediaFileBitrateTooHigh = errors.New("The bitrate of the MediaFile is too high")

	ErrMediaFileBitrateTooLow = errors.New("The bitrate of the MediaFile is too low")

	ErrMediaFileHeightTooHigh = errors.New("The height of the MediaFile is too high")

	ErrMediaFileHeightTooLow = errors.New("The height of the MediaFile is too low")

	ErrMediaFileMissDelivery = errors.New("ErrMediaFile miss Delivery")

	ErrMediaFileMissMimeType = errors.New("ErrMediaFile miss MimeType")

	ErrMediaFileMissUri = errors.New("ErrMediaFile miss Uri")

	ErrMediaFileSize = errors.New("The width and height of MediaFile should be greater than zero")

	ErrMediaFileUnsupportedMimeType = errors.New("The MimeType in the MediaFile is not supported")

	ErrMediaFileWidthTooHigh = errors.New("The width of the MediaFile is too high")

	ErrMediaFileWidthTooLow = errors.New("The width of the MediaFile is too low")

	ErrNonLinearAdsMissNonLinears = errors.New("NonLinearAds miss NonLinears")

	ErrNonLinearResourceFormat = errors.New("NonLinear resources should contain only one of StaticResource, IFrameResource, HtmlResource")

	ErrOffsetPercentNegative = errors.New("Percent in Offset is neagtive")

	ErrPricingCurrencyFormat = errors.New("Pricing's Currency must comply ISO-4217 3 letter format")

	ErrPricingMissModel = errors.New("Pricing misses Model")

	ErrPricingMissPrice = errors.New("Pricing misses Price")

	ErrUnsupportedDeliveryType = errors.New("Delivery does not equal `progressive` and `streaming`")

	ErrUnsupportedEvent = errors.New("Event type is unsupported")

	ErrUnsupportedMode = errors.New("Mode is not supported")

	ErrUnsupportedPriceModelType = errors.New("PriceModelType is unsuppored")

	ErrUnsupportedVersion = errors.New("Version is not supported")

	ErrVastMissAd = errors.New("Vast misses Ad")

	ErrVideoClicksMissClickThroughs = errors.New("ClickThroughs are missed")

	ErrVideoDurationTooLong = errors.New("Video duration is too long")

	ErrVideoDurationTooShort = errors.New("Video duration is too short")

	ErrWrapperMissAdSystem = errors.New("Wrapper misses AdSystem")

	ErrWrapperMissImpressions = errors.New("Wrapper misses Impressions")

	ErrWrapperMissVastAdTagUri = errors.New("Wrapper misses VastAdTagUri")
)
