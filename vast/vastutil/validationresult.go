package vastutil

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/vast2"
	"net/http"
	"strconv"
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

	case vast2.ErrAdSystemMissSystem:
		return ResultAdSystemMissSystem
	case vast2.ErrAdType:
		return ResultAdType
	case vast2.ErrCompanionAdsMissCompanions:
		return ResultCompanionAdsMissCompanions
	case vast2.ErrCompanionAdsWrapperMissCompanions:
		return ResultCompanionAdsWrapperMissCompanions
	case vast2.ErrCreativeType:
		return ResultCreativeType
	case vast2.ErrCreativeWrapperType:
		return ResultCreativeWrapperType
	case vast2.ErrDurationEqualZero:
		return ResultDurationEqualZero
	case vast2.ErrDurationNegative:
		return ResultDurationNegative
	case vast2.ErrIconMissPosition:
		return ResultIconMissPosition
	case vast2.ErrIconMissProgram:
		return ResultIconMissProgram
	case vast2.ErrIconResourcesFormat:
		return ResultIconResourcesFormat
	case vast2.ErrImpressionMissURI:
		return ResultImpressionMissURI
	case vast2.ErrInlineMissAdTitle:
		return ResultInlineMissAdTitle
	case vast2.ErrInlineMissCreatives:
		return ResultInlineMissCreatives
	case vast2.ErrInlineMissImpressions:
		return ResultInlineMissImpressions
	case vast2.ErrLinearMissMediaFiles:
		return ResultLinearMissMediaFiles
	case vast2.ErrMediaFileHeightTooHigh:
		return ResultMediaFileHeightTooHigh
	case vast2.ErrMediaFileHeightTooLow:
		return ResultModiaFileHeightTooLow
	case vast2.ErrMediaFileMissDelivery:
		return ResultMediaFileMissDelivery
	case vast2.ErrMediaFileMissMimeType:
		return ResultMediaFileMissMimeType
	case vast2.ErrMediaFileMissURI:
		return ResultMediaFileMissURI
	case vast2.ErrMediaFileSize:
		return ResultMediaFileSize
	case vast2.ErrMediaFileUnsupportedMimeType:
		return ResultMediaFileUnsupportedMimeType
	case vast2.ErrMediaFileWidthTooHigh:
		return ResultMediaFileWidthTooHigh
	case vast2.ErrMediaFileWidthTooLow:
		return ResultMediaFileWidthTooLow
	case vast2.ErrNonLinearAdsMissNonLinears:
		return ResultNonLinearAdsMissNonLinears
	case vast2.ErrNonLinearResourceFormat:
		return ResultNonLinearResourceFormat
	case vast2.ErrOffsetPercentNegative:
		return ResultOffsetPercentNegative
	case vast2.ErrPricingCurrencyFormat:
		return ResultPricingCurrencyFormat
	case vast2.ErrPricingMissModel:
		return ResultPricingMissModel
	case vast2.ErrPricingMissPrice:
		return ResultPricingMissPrice
	case vast2.ErrUnsupportedDeliveryType:
		return ResultUnsupportedDeliveryType
	case vast2.ErrUnsupportedEvent:
		return ResultUnsupportedEvent
	case vast2.ErrUnsupportedMode:
		return ResultUnsupportedMode
	case vast2.ErrUnsupportedPriceModelType:
		return ResultUnsupportedPriceModelType
	case vast2.ErrUnsupportedVersion:
		return ResultUnsupportedVersion
	case vast2.ErrVastMissAd:
		return ResultVastMissAd
	case vast2.ErrVideoClicksMissClickThroughs:
		return ResultVideoClicksMissClickThroughs
	case vast2.ErrVideoDurationTooLong:
		return ResultVideoDurationTooLong
	case vast2.ErrVideoDurationTooShort:
		return ResultVideoDurationTooShort
	case vast2.ErrWrapperMissImpressions:
		return ResultWrapperMissImpressions
	case vast2.ErrWrapperMissVastAdTagURI:
		return ResultWrapperMissVastAdTagURI
	}

	return ResultUnknown
}

// EncodeInvalidVastHeader method encodes the invalid vast result into a string header value in the HTTP header.
func EncodeInvalidVastHeader(header http.Header, vastErr error) {
	validationResult := getValidationResultFromErr(vastErr)
	header.Set(HeaderXInvalidVastResult, strconv.Itoa(int(validationResult)))
}
