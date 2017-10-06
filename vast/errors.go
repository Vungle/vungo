package vast

import "errors"

// Please keep the variables alphabetized.

var ErrAdParametersMissParameters = errors.New("AdParameters misses Parameters.")

var ErrAdSystemMissSystem = errors.New("AdSystem misses system.")

var ErrAdType = errors.New("Ad should only contain one of Inline and Wrapper.")

var ErrCompanionAdsMissCompanions = errors.New("CompanionAds miss Companions.")

var ErrCompanionAdsWrapperMissCompanions = errors.New("CompanionAdsWrapper misses Companions.")

var ErrCompanionResourceFormat = errors.New("Companion should only contain one of IFrameResource, StaticResource, and HtmlResource.")

var ErrCompanionWrapperResourceFormat = errors.New("CompanionWrapperResource should only contain one of IFrameResource, StaticResource, and HtmlResource.")

var ErrCreativeType = errors.New("Creative should contain one type of NonLinearAds, Linear, CompanionAds.")

var ErrCreativeWrapperType = errors.New("CreativeWrapper should contain one type of NonLinearAds, Linear, CompanionAds.")

var ErrDurationEqualZero = errors.New("Duration is less than zero.")

var ErrDurationNegative = errors.New("Duration is negative.")

var ErrHtmlResourceMissHtml = errors.New("HtmlResource misses Html.")

var ErrIconMissPosition = errors.New("Icon misses position.")

var ErrIconMissProgram = errors.New("Icon misses program.")

var ErrIconResourcesFormat = errors.New("Icon resources should contain only one of StaticResource, IFrameResource, HtmlResource.")

var ErrImpressionMissUri = errors.New("Impression misses Uri.")

var ErrInlineMissAdTitle = errors.New("Inline misses AdTitle.")

var ErrInlineMissCreatives = errors.New("Inline AD misses Creatives.")

var ErrInlineMissImpressions = errors.New("Inline misses Impressions.")

var ErrLinearMissMediaFiles = errors.New("MediaFiles are missed.")

var ErrMediaFileBitrateTooHigh = errors.New("The bitrate of the MediaFile is too high.")

var ErrMediaFileBitrateTooLow = errors.New("The bitrate of the MediaFile is too low.")

var ErrMediaFileHeightTooHigh = errors.New("The height of the MediaFile is too high.")

var ErrMediaFileHeightTooLow = errors.New("The height of the MediaFile is too low.")

var ErrMediaFileMissDelivery = errors.New("ErrMediaFile miss Delivery.")

var ErrMediaFileMissMimeType = errors.New("ErrMediaFile miss MimeType.")

var ErrMediaFileMissUri = errors.New("ErrMediaFile miss Uri.")

var ErrMediaFileSize = errors.New("The width and height of MediaFile should be greater than zero.")

var ErrMediaFileUnsupportedMimeType = errors.New("The MimeType in the MediaFile is not supported.")

var ErrMediaFileWidthTooHigh = errors.New("The width of the MediaFile is too high.")

var ErrMediaFileWidthTooLow = errors.New("The width of the MediaFile is too low.")

var ErrNonLinearAdsMissNonLinears = errors.New("NonLinearAds miss NonLinears.")

var ErrNonLinearResourceFormat = errors.New("NonLinear resources should contain only one of StaticResource, IFrameResource, HtmlResource.")

var ErrOffsetPercentNegative = errors.New("Percent in Offset is neagtive.")

var ErrPricingCurrencyFormat = errors.New("Pricing's Currency must comply ISO-4217 3 letter format.")

var ErrPricingMissModel = errors.New("Pricing misses Model.")

var ErrPricingMissPrice = errors.New("Pricing misses Price.")

var ErrUnsupportedDeliveryType = errors.New("Delivery does not equal `progressive` and `streaming`.")

var ErrUnsupportedEvent = errors.New("Event type is unsupported.")

var ErrUnsupportedMode = errors.New("Mode is not supported.")

var ErrUnsupportedPriceModelType = errors.New("PriceModelType is unsuppored.")

var ErrUnsupportedVersion = errors.New("Version is not supported.")

var ErrVastMissAd = errors.New("Vast misses Ad.")

var ErrVideoClicksMissClickThroughs = errors.New("ClickThroughs are missed.")

var ErrVideoDurationTooLong = errors.New("Video duration is too long.")

var ErrVideoDurationTooShort = errors.New("Video duration is too short.")

var ErrWrapperMissImpressions = errors.New("Wrapper misses Impressions.")

var ErrWrapperMissVastAdTagUri = errors.New("Wrapper misses VastAdTagUri.")
