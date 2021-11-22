package vast

import "errors"

// Please keep the variables alphabetized.
var (
	// ErrAdSystemMissSystem error for AdSystem misses system.
	ErrAdSystemMissSystem = errors.New("AdSystem misses system")

	// ErrAdType error for Ad should only contain one of Inline and Wrapper.
	ErrAdType = errors.New("Ad should only contain one of Inline and Wrapper")

	// ErrCompanionAdsMissCompanions error for CompanionAds miss Companions.
	ErrCompanionAdsMissCompanions = errors.New("CompanionAds miss Companions")

	// ErrCompanionAdsWrapperMissCompanions error for CompanionAdsWrapper misses Companions.
	ErrCompanionAdsWrapperMissCompanions = errors.New("CompanionAdsWrapper misses Companions")

	// 	ErrCreativeType error for Creative should contain one type of NonLinearAds, Linear, CompanionAds.
	ErrCreativeType = errors.New("Creative should contain one type of NonLinearAds, Linear, CompanionAds")

	// ErrCreativeWrapperType error for CreativeWrapper should contain one type of NonLinearAds, Linear, CompanionAds.
	ErrCreativeWrapperType = errors.New("CreativeWrapper should contain one type of NonLinearAds, Linear, CompanionAds")

	// ErrDurationEqualZero error for Duration is less than zero.
	ErrDurationEqualZero = errors.New("Duration is less than zero")

	// ErrDurationNegative error for Duration is negative.
	ErrDurationNegative = errors.New("Duration is negative")

	// ErrIconMissPosition error for Icon misses position.
	ErrIconMissPosition = errors.New("Icon misses position")

	// ErrIconMissProgram error for Icon misses program.
	ErrIconMissProgram = errors.New("Icon misses program")

	// ErrIconResourcesFormat error for Icon resources should contain only one of StaticResource, IFrameResource, HTMLResource.
	ErrIconResourcesFormat = errors.New("Icon resources should contain only one of StaticResource, IFrameResource, HTMLResource")

	// ErrImpressionMissURI error for Impression misses URI.
	ErrImpressionMissURI = errors.New("Impression misses URI")

	// ErrInlineMissAdSystem error for Inline misses AdSystem.
	ErrInlineMissAdSystem = errors.New("Inline misses AdSystem")

	// ErrInlineMissAdTitle error for Inline misses AdTitle.
	ErrInlineMissAdTitle = errors.New("Inline misses AdTitle")

	// ErrInlineMissCreatives error for Inline AD misses Creatives.
	ErrInlineMissCreatives = errors.New("Inline AD misses Creatives")

	// ErrInlineMissImpressions error for Inline misses Impressions.
	ErrInlineMissImpressions = errors.New("Inline misses Impressions")

	// ErrLinearMissMediaFiles error for MediaFiles are missed.
	ErrLinearMissMediaFiles = errors.New("MediaFiles are missed")

	// ErrMediaFileHeightTooHigh error for The height of the MediaFile is too high.
	ErrMediaFileHeightTooHigh = errors.New("The height of the MediaFile is too high")

	// ErrMediaFileHeightTooLow error for The height of the MediaFile is too low.
	ErrMediaFileHeightTooLow = errors.New("The height of the MediaFile is too low")

	// ErrMediaFileMissDelivery error for ErrMediaFile miss Delivery.
	ErrMediaFileMissDelivery = errors.New("ErrMediaFile miss Delivery")

	// ErrMediaFileMissMimeType error for ErrMediaFile miss MimeType.
	ErrMediaFileMissMimeType = errors.New("ErrMediaFile miss MimeType")

	// ErrMediaFileMissURI error for ErrMediaFile miss URI.
	ErrMediaFileMissURI = errors.New("ErrMediaFile miss URI")

	// ErrMediaFileSize The width and height of MediaFile should be greater than zero.
	ErrMediaFileSize = errors.New("The width and height of MediaFile should be greater than zero")

	// ErrMediaFileUnsupportedMimeType The MimeType in the MediaFile is not supported.
	ErrMediaFileUnsupportedMimeType = errors.New("The MimeType in the MediaFile is not supported")

	// ErrMediaFileWidthTooHigh error for The width of the MediaFile is too high.
	ErrMediaFileWidthTooHigh = errors.New("The width of the MediaFile is too high")

	// ErrMediaFileWidthTooLow error for The width of the MediaFile is too low.
	ErrMediaFileWidthTooLow = errors.New("The width of the MediaFile is too low")

	// ErrMissAdInline error for Ad should only contain one of Inline and Wrapper.
	ErrMissAdInline = errors.New("Ad should only contain one of Inline and Wrapper")

	// ErrNonLinearAdsMissNonLinears error for NonLinearAds miss NonLinears.
	ErrNonLinearAdsMissNonLinears = errors.New("NonLinearAds miss NonLinears")

	// ErrNonLinearResourceFormat error for NonLinear resources should contain only one of StaticResource, IFrameResource, HTMLResource.
	ErrNonLinearResourceFormat = errors.New("NonLinear resources should contain only one of StaticResource, IFrameResource, HTMLResource")

	// ErrOffsetPercentNegative error for Percent in Offset is neagtive.
	ErrOffsetPercentNegative = errors.New("Percent in Offset is neagtive")

	// ErrPricingCurrencyFormat error for Pricing's Currency must comply ISO-4217 3 letter format.
	ErrPricingCurrencyFormat = errors.New("Pricing's Currency must comply ISO-4217 3 letter format")

	// ErrPricingMissModel error for Pricing misses Model.
	ErrPricingMissModel = errors.New("Pricing misses Model")

	// ErrPricingMissPrice error for Pricing misses Price.
	ErrPricingMissPrice = errors.New("Pricing misses Price")

	// ErrUnsupportedDeliveryType error for Delivery does not equal `progressive` and `streaming`.
	ErrUnsupportedDeliveryType = errors.New("Delivery does not equal `progressive` and `streaming`")

	// ErrUnsupportedEvent error for Event type is unsupported.
	ErrUnsupportedEvent = errors.New("Event type is unsupported")

	// ErrUnsupportedMode error for Mode is not supported.
	ErrUnsupportedMode = errors.New("Mode is not supported")

	// ErrUnsupportedPriceModelType error for PriceModelType is unsuppored.
	ErrUnsupportedPriceModelType = errors.New("PriceModelType is unsuppored")

	// ErrUnsupportedVersion error for Version is not supported.
	ErrUnsupportedVersion = errors.New("Version is not supported")

	// ErrVastMissAd error for Vast misses Ad.
	ErrVastMissAd = errors.New("Vast misses Ad")

	// ErrVideoClicksMissClickThroughs error for lickThroughs are missed.
	ErrVideoClicksMissClickThroughs = errors.New("ClickThroughs are missed")

	// ErrVideoDurationTooLong error for Video duration is too long.
	ErrVideoDurationTooLong = errors.New("Video duration is too long")

	// ErrVideoDurationTooShort error for Video duration is too short.
	ErrVideoDurationTooShort = errors.New("Video duration is too short")

	// ErrWrapperMissAdSystem error for Wrapper misses AdSystem.
	ErrWrapperMissAdSystem = errors.New("Wrapper misses AdSystem")

	// ErrWrapperMissImpressions error for Wrapper misses Impressions.
	ErrWrapperMissImpressions = errors.New("Wrapper misses Impressions")

	// ErrWrapperMissVastAdTagURI error for Wrapper misses VastAdTagURI.
	ErrWrapperMissVastAdTagURI = errors.New("Wrapper misses VastAdTagURI")
)
