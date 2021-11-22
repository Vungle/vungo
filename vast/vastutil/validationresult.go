package vastutil

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/Vungle/vungo/vast"
)

// VastValidationResult type represents the result of unwrapping and validating VAST objects.
type VastValidationResult int

// HeaderXInvalidVastResult value is an HTTP header that will be embedded in an ad response to
// indicate an invalid VAST result for a particular ad request.
const HeaderXInvalidVastResult = "X-Vungle-Invalid-Vast"

// The following constants are an enumeration of all possible results output by the VAST unwrapping &
// validation process.
const (
	ResultUnknown VastValidationResult = -1
	ResultOK      VastValidationResult = 0

	ResultXMLSyntaxError          VastValidationResult = 1
	ResultXMLTagPathError         VastValidationResult = 2
	ResultXMLUnmarshalError       VastValidationResult = 3
	ResultXMLUnsupportedTypeError VastValidationResult = 4

	ResultUnwrapWithMultipleAds  VastValidationResult = 10000
	ResultWrapperMissingAdTagURI VastValidationResult = 10001

	ResultAdSystemMissSystem                VastValidationResult = 20001
	ResultAdType                            VastValidationResult = 20002
	ResultCompanionAdsMissCompanions        VastValidationResult = 20003
	ResultCompanionAdsWrapperMissCompanions VastValidationResult = 20004
	ResultCompanionResourceFormat           VastValidationResult = 20005
	ResultCompanionWrapperResourceFormat    VastValidationResult = 20006
	ResultCreativeType                      VastValidationResult = 20007
	ResultCreativeWrapperType               VastValidationResult = 20008
	ResultDurationEqualZero                 VastValidationResult = 20009
	ResultDurationNegative                  VastValidationResult = 20010
	ResultHTMLResourceMissHTML              VastValidationResult = 20011
	ResultIconMissPosition                  VastValidationResult = 20012
	ResultIconMissProgram                   VastValidationResult = 20013
	ResultIconResourcesFormat               VastValidationResult = 20014
	ResultInlineMissAdTitle                 VastValidationResult = 20015
	ResultInlineMissCreatives               VastValidationResult = 20016
	ResultInlineMissImpressions             VastValidationResult = 20017
	ResultImpressionMissURI                 VastValidationResult = 20018
	ResultLinearMissMediaFiles              VastValidationResult = 20019
	ResultMediaFileMissDelivery             VastValidationResult = 20020
	ResultMediaFileMissMimeType             VastValidationResult = 20021
	ResultMediaFileMissURI                  VastValidationResult = 20022
	ResultMediaFileSize                     VastValidationResult = 20023
	ResultNonLinearAdsMissNonLinears        VastValidationResult = 20024
	ResultNonLinearResourceFormat           VastValidationResult = 20025
	ResultOffsetPercentNegative             VastValidationResult = 20026
	ResultPricingCurrencyFormat             VastValidationResult = 20027
	ResultPricingMissModel                  VastValidationResult = 20028
	ResultPricingMissPrice                  VastValidationResult = 20029
	ResultUnsupportedDeliveryType           VastValidationResult = 20032
	ResultUnsupportedEvent                  VastValidationResult = 20033
	ResultUnsupportedMode                   VastValidationResult = 20034
	ResultUnsupportedPriceModelType         VastValidationResult = 20035
	ResultUnsupportedVersion                VastValidationResult = 20036
	ResultVastMissAd                        VastValidationResult = 20037
	ResultVideoClicksMissClickThroughs      VastValidationResult = 20039
	ResultWrapperMissImpressions            VastValidationResult = 20040
	ResultWrapperMissVastAdTagURI           VastValidationResult = 20041
	ResultVideoDurationTooShort             VastValidationResult = 20042
	ResultVideoDurationTooLong              VastValidationResult = 20043
	ResultMediaFileHeightTooHigh            VastValidationResult = 20046
	ResultModiaFileHeightTooLow             VastValidationResult = 20047
	ResultMediaFileUnsupportedMimeType      VastValidationResult = 20048
	ResultMediaFileWidthTooHigh             VastValidationResult = 20049
	ResultMediaFileWidthTooLow              VastValidationResult = 20050
)

func getValidationResultFromErr(err error) VastValidationResult {
	switch err.(type) {
	case nil:
		return ResultOK
	case *xml.SyntaxError:
		return ResultXMLSyntaxError
	case *xml.TagPathError:
		return ResultXMLTagPathError
	case *xml.UnmarshalError:
		return ResultXMLUnmarshalError
	case *xml.UnsupportedTypeError:
		return ResultXMLUnsupportedTypeError
	}

	switch err {
	case ErrUnwrapWithMultipleAds:
		return ResultUnwrapWithMultipleAds
	case ErrWrapperMissingAdTagURI:
		return ResultWrapperMissingAdTagURI

	case vast.ErrAdSystemMissSystem:
		return ResultAdSystemMissSystem
	case vast.ErrAdType:
		return ResultAdType
	case vast.ErrCompanionAdsMissCompanions:
		return ResultCompanionAdsMissCompanions
	case vast.ErrCompanionAdsWrapperMissCompanions:
		return ResultCompanionAdsWrapperMissCompanions
	case vast.ErrCreativeType:
		return ResultCreativeType
	case vast.ErrCreativeWrapperType:
		return ResultCreativeWrapperType
	case vast.ErrDurationEqualZero:
		return ResultDurationEqualZero
	case vast.ErrDurationNegative:
		return ResultDurationNegative
	case vast.ErrIconMissPosition:
		return ResultIconMissPosition
	case vast.ErrIconMissProgram:
		return ResultIconMissProgram
	case vast.ErrIconResourcesFormat:
		return ResultIconResourcesFormat
	case vast.ErrImpressionMissURI:
		return ResultImpressionMissURI
	case vast.ErrInlineMissAdTitle:
		return ResultInlineMissAdTitle
	case vast.ErrInlineMissCreatives:
		return ResultInlineMissCreatives
	case vast.ErrInlineMissImpressions:
		return ResultInlineMissImpressions
	case vast.ErrLinearMissMediaFiles:
		return ResultLinearMissMediaFiles
	case vast.ErrMediaFileHeightTooHigh:
		return ResultMediaFileHeightTooHigh
	case vast.ErrMediaFileHeightTooLow:
		return ResultModiaFileHeightTooLow
	case vast.ErrMediaFileMissDelivery:
		return ResultMediaFileMissDelivery
	case vast.ErrMediaFileMissMimeType:
		return ResultMediaFileMissMimeType
	case vast.ErrMediaFileMissURI:
		return ResultMediaFileMissURI
	case vast.ErrMediaFileSize:
		return ResultMediaFileSize
	case vast.ErrMediaFileUnsupportedMimeType:
		return ResultMediaFileUnsupportedMimeType
	case vast.ErrMediaFileWidthTooHigh:
		return ResultMediaFileWidthTooHigh
	case vast.ErrMediaFileWidthTooLow:
		return ResultMediaFileWidthTooLow
	case vast.ErrNonLinearAdsMissNonLinears:
		return ResultNonLinearAdsMissNonLinears
	case vast.ErrNonLinearResourceFormat:
		return ResultNonLinearResourceFormat
	case vast.ErrOffsetPercentNegative:
		return ResultOffsetPercentNegative
	case vast.ErrPricingCurrencyFormat:
		return ResultPricingCurrencyFormat
	case vast.ErrPricingMissModel:
		return ResultPricingMissModel
	case vast.ErrPricingMissPrice:
		return ResultPricingMissPrice
	case vast.ErrUnsupportedDeliveryType:
		return ResultUnsupportedDeliveryType
	case vast.ErrUnsupportedEvent:
		return ResultUnsupportedEvent
	case vast.ErrUnsupportedMode:
		return ResultUnsupportedMode
	case vast.ErrUnsupportedPriceModelType:
		return ResultUnsupportedPriceModelType
	case vast.ErrUnsupportedVersion:
		return ResultUnsupportedVersion
	case vast.ErrVastMissAd:
		return ResultVastMissAd
	case vast.ErrVideoClicksMissClickThroughs:
		return ResultVideoClicksMissClickThroughs
	case vast.ErrVideoDurationTooLong:
		return ResultVideoDurationTooLong
	case vast.ErrVideoDurationTooShort:
		return ResultVideoDurationTooShort
	case vast.ErrWrapperMissImpressions:
		return ResultWrapperMissImpressions
	case vast.ErrWrapperMissVastAdTagURI:
		return ResultWrapperMissVastAdTagURI
	}

	return ResultUnknown
}

// EncodeInvalidVastHeader method encodes the invalid vast result into a string header value in the HTTP header.
func EncodeInvalidVastHeader(header http.Header, vastErr error) {
	validationResult := getValidationResultFromErr(vastErr)
	header.Set(HeaderXInvalidVastResult, strconv.Itoa(int(validationResult)))
}
