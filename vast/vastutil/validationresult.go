package vastutil

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/Vungle/vungo/vast"
)

// VastValidationResult type represents the result of unwrapping and validating VAST objects.
type VastValidationResult int

// HEADER_X_INVALID_VAST_RESULT value is an HTTP header that will be embedded in an ad response to
// indicate an invalid VAST result for a particular ad request.
const HEADER_X_INVALID_VAST_RESULT = "X-Vungle-Invalid-Vast"

// The following constants are an enumeration of all possible results output by the VAST unwrapping &
// validation process.
const (
	RESULT_UNKNOWN VastValidationResult = -1

	RESULT_OK VastValidationResult = 0

	RESULT_XML_SYNTAX_ERROR           VastValidationResult = 1
	RESULT_XML_TAG_PATH_ERROR         VastValidationResult = 2
	RESULT_XML_UNMARSHAL_ERROR        VastValidationResult = 3
	RESULT_XML_UNSUPPORTED_TYPE_ERROR VastValidationResult = 4

	RESULT_UNWRAP_WITH_MULTIPLE_ADS   VastValidationResult = 10000
	RESULT_WRAPPER_MISSING_AD_TAG_URI VastValidationResult = 10001

	RESULT_AD_PARAMETERS_MISS_PARAMETERS         VastValidationResult = 20000
	RESULT_AD_SYSTEM_MISS_SYSTEM                 VastValidationResult = 20001
	RESULT_AD_TYPE                               VastValidationResult = 20002
	RESULT_COMPANION_ADS_MISS_COMPANIONS         VastValidationResult = 20003
	RESULT_COMPANION_ADS_WRAPPER_MISS_COMPANIONS VastValidationResult = 20004
	RESULT_COMPANION_RESOURCE_FORMAT             VastValidationResult = 20005
	RESULT_COMPANION_WRAPPER_RESOURCE_FORMAT     VastValidationResult = 20006
	RESULT_CREATIVE_TYPE                         VastValidationResult = 20007
	RESULT_CREATIVE_WRAPPER_TYPE                 VastValidationResult = 20008
	RESULT_DURATION_EQUAL_ZERO                   VastValidationResult = 20009
	RESULT_DURATION_NEGATIVE                     VastValidationResult = 20010
	RESULT_HTML_RESOURCE_MISS_HTML               VastValidationResult = 20011
	RESULT_ICON_MISS_POSITION                    VastValidationResult = 20012
	RESULT_ICON_MISS_PROGRAM                     VastValidationResult = 20013
	RESULT_ICON_RESOURCES_FORMAT                 VastValidationResult = 20014
	RESULT_INLINE_MISS_AD_TITLE                  VastValidationResult = 20015
	RESULT_INLINE_MISS_CREATIVES                 VastValidationResult = 20016
	RESULT_INLINE_MISS_IMPRESSIONS               VastValidationResult = 20017
	RESULT_IMPRESSION_MISS_URI                   VastValidationResult = 20018
	RESULT_LINEAR_MISS_MEDIA_FILES               VastValidationResult = 20019
	RESULT_MEDIA_FILE_MISS_DELIVERY              VastValidationResult = 20020
	RESULT_MEDIA_FILE_MISS_MIME_TYPE             VastValidationResult = 20021
	RESULT_MEDIA_FILE_MISS_URI                   VastValidationResult = 20022
	RESULT_MEDIA_FILE_SIZE                       VastValidationResult = 20023
	RESULT_NON_LINEAR_ADS_MISS_NON_LINEARS       VastValidationResult = 20024
	RESULT_NON_LINEAR_RESOURCE_FORMAT            VastValidationResult = 20025
	RESULT_OFFSET_PERCENT_NEGATIVE               VastValidationResult = 20026
	RESULT_PRICING_CURRENCY_FORMAT               VastValidationResult = 20027
	RESULT_PRICING_MISS_MODEL                    VastValidationResult = 20028
	RESULT_PRICING_MISS_PRICE                    VastValidationResult = 20029
	RESULT_UNSUPPORTED_DELIVERY_TYPE             VastValidationResult = 20032
	RESULT_UNSUPPORTED_EVENT                     VastValidationResult = 20033
	RESULT_UNSUPPORTED_MODE                      VastValidationResult = 20034
	RESULT_UNSUPPORTED_PRICE_MODEL_TYPE          VastValidationResult = 20035
	RESULT_UNSUPPORTED_VERSION                   VastValidationResult = 20036
	RESULT_VAST_MISS_AD                          VastValidationResult = 20037
	RESULT_VIDEO_CLICKS_MISS_CLICK_THROUGHS      VastValidationResult = 20039
	RESULT_WRAPPER_MISS_IMPRESSIONS              VastValidationResult = 20040
	RESULT_WRAPPER_MISS_VAST_AD_TAG_URI          VastValidationResult = 20041
	RESULT_VIDEO_DURATION_TOO_SHORT              VastValidationResult = 20042
	RESULT_VIDEO_DURATION_TOO_LONG               VastValidationResult = 20043
	RESULT_MEDIA_FILE_BITRATE_TOO_HIGH           VastValidationResult = 20044
	RESULT_MEDIA_FILE_BITRATE_TOO_LOW            VastValidationResult = 20045
	RESULT_MEDIA_FILE_HEIGHT_TOO_HIGH            VastValidationResult = 20046
	RESULT_MEDIA_FILE_HEIGHT_TOO_LOW             VastValidationResult = 20047
	RESULT_MEDIA_FILE_UNSUPPORTED_MIME_TYPE      VastValidationResult = 20048
	RESULT_MEDIA_FILE_WIDTH_TOO_HIGH             VastValidationResult = 20049
	RESULT_MEDIA_FILE_WIDTH_TOO_LOW              VastValidationResult = 20050
)

func getValidationResultFromErr(err error) VastValidationResult {
	switch err.(type) {
	case nil:
		return RESULT_OK
	case *xml.SyntaxError:
		return RESULT_XML_SYNTAX_ERROR
	case *xml.TagPathError:
		return RESULT_XML_TAG_PATH_ERROR
	case *xml.UnmarshalError:
		return RESULT_XML_UNMARSHAL_ERROR
	case *xml.UnsupportedTypeError:
		return RESULT_XML_UNSUPPORTED_TYPE_ERROR
	}

	switch err {
	case ErrUnwrapWithMultipleAds:
		return RESULT_UNWRAP_WITH_MULTIPLE_ADS
	case ErrWrapperMissingAdTagUri:
		return RESULT_WRAPPER_MISSING_AD_TAG_URI

	case vast.ErrAdParametersMissParameters:
		return RESULT_AD_PARAMETERS_MISS_PARAMETERS
	case vast.ErrAdSystemMissSystem:
		return RESULT_AD_SYSTEM_MISS_SYSTEM
	case vast.ErrAdType:
		return RESULT_AD_TYPE
	case vast.ErrCompanionAdsMissCompanions:
		return RESULT_COMPANION_ADS_MISS_COMPANIONS
	case vast.ErrCompanionAdsWrapperMissCompanions:
		return RESULT_COMPANION_ADS_WRAPPER_MISS_COMPANIONS
	case vast.ErrCompanionResourceFormat:
		return RESULT_COMPANION_RESOURCE_FORMAT
	case vast.ErrCompanionWrapperResourceFormat:
		return RESULT_COMPANION_WRAPPER_RESOURCE_FORMAT
	case vast.ErrCreativeType:
		return RESULT_CREATIVE_TYPE
	case vast.ErrCreativeWrapperType:
		return RESULT_CREATIVE_WRAPPER_TYPE
	case vast.ErrDurationEqualZero:
		return RESULT_DURATION_EQUAL_ZERO
	case vast.ErrDurationNegative:
		return RESULT_DURATION_NEGATIVE
	case vast.ErrHtmlResourceMissHtml:
		return RESULT_HTML_RESOURCE_MISS_HTML
	case vast.ErrIconMissPosition:
		return RESULT_ICON_MISS_POSITION
	case vast.ErrIconMissProgram:
		return RESULT_ICON_MISS_PROGRAM
	case vast.ErrIconResourcesFormat:
		return RESULT_ICON_RESOURCES_FORMAT
	case vast.ErrImpressionMissUri:
		return RESULT_IMPRESSION_MISS_URI
	case vast.ErrInlineMissAdTitle:
		return RESULT_INLINE_MISS_AD_TITLE
	case vast.ErrInlineMissCreatives:
		return RESULT_INLINE_MISS_CREATIVES
	case vast.ErrInlineMissImpressions:
		return RESULT_INLINE_MISS_IMPRESSIONS
	case vast.ErrLinearMissMediaFiles:
		return RESULT_LINEAR_MISS_MEDIA_FILES
	case vast.ErrMediaFileBitrateTooHigh:
		return RESULT_MEDIA_FILE_BITRATE_TOO_HIGH
	case vast.ErrMediaFileBitrateTooLow:
		return RESULT_MEDIA_FILE_BITRATE_TOO_LOW
	case vast.ErrMediaFileHeightTooHigh:
		return RESULT_MEDIA_FILE_HEIGHT_TOO_HIGH
	case vast.ErrMediaFileHeightTooLow:
		return RESULT_MEDIA_FILE_HEIGHT_TOO_LOW
	case vast.ErrMediaFileMissDelivery:
		return RESULT_MEDIA_FILE_MISS_DELIVERY
	case vast.ErrMediaFileMissMimeType:
		return RESULT_MEDIA_FILE_MISS_MIME_TYPE
	case vast.ErrMediaFileMissUri:
		return RESULT_MEDIA_FILE_MISS_URI
	case vast.ErrMediaFileSize:
		return RESULT_MEDIA_FILE_SIZE
	case vast.ErrMediaFileUnsupportedMimeType:
		return RESULT_MEDIA_FILE_UNSUPPORTED_MIME_TYPE
	case vast.ErrMediaFileWidthTooHigh:
		return RESULT_MEDIA_FILE_WIDTH_TOO_HIGH
	case vast.ErrMediaFileWidthTooLow:
		return RESULT_MEDIA_FILE_WIDTH_TOO_LOW
	case vast.ErrNonLinearAdsMissNonLinears:
		return RESULT_NON_LINEAR_ADS_MISS_NON_LINEARS
	case vast.ErrNonLinearResourceFormat:
		return RESULT_NON_LINEAR_RESOURCE_FORMAT
	case vast.ErrOffsetPercentNegative:
		return RESULT_OFFSET_PERCENT_NEGATIVE
	case vast.ErrPricingCurrencyFormat:
		return RESULT_PRICING_CURRENCY_FORMAT
	case vast.ErrPricingMissModel:
		return RESULT_PRICING_MISS_MODEL
	case vast.ErrPricingMissPrice:
		return RESULT_PRICING_MISS_PRICE
	case vast.ErrUnsupportedDeliveryType:
		return RESULT_UNSUPPORTED_DELIVERY_TYPE
	case vast.ErrUnsupportedEvent:
		return RESULT_UNSUPPORTED_EVENT
	case vast.ErrUnsupportedMode:
		return RESULT_UNSUPPORTED_MODE
	case vast.ErrUnsupportedPriceModelType:
		return RESULT_UNSUPPORTED_PRICE_MODEL_TYPE
	case vast.ErrUnsupportedVersion:
		return RESULT_UNSUPPORTED_VERSION
	case vast.ErrVastMissAd:
		return RESULT_VAST_MISS_AD
	case vast.ErrVideoClicksMissClickThroughs:
		return RESULT_VIDEO_CLICKS_MISS_CLICK_THROUGHS
	case vast.ErrVideoDurationTooLong:
		return RESULT_VIDEO_DURATION_TOO_LONG
	case vast.ErrVideoDurationTooShort:
		return RESULT_VIDEO_DURATION_TOO_SHORT
	case vast.ErrWrapperMissImpressions:
		return RESULT_WRAPPER_MISS_IMPRESSIONS
	case vast.ErrWrapperMissVastAdTagUri:
		return RESULT_WRAPPER_MISS_VAST_AD_TAG_URI
	}

	return RESULT_UNKNOWN
}

// EncodeInvalidVastHeader method encodes the invalid vast result into a string header value in the HTTP header.
func EncodeInvalidVastHeader(header http.Header, vastErr error) {
	validationResult := getValidationResultFromErr(vastErr)
	header.Set(HEADER_X_INVALID_VAST_RESULT, strconv.Itoa(int(validationResult)))
}
