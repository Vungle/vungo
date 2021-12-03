package vastutil

import (
	"encoding/xml"
	"github.com/Vungle/vungo/vast/basic"
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

	case vastbasic.ErrAdSystemMissSystem:
		return ResultAdSystemMissSystem
	case vastbasic.ErrAdType:
		return ResultAdType
	case vastbasic.ErrCompanionAdsMissCompanions:
		return ResultCompanionAdsMissCompanions
	case vastbasic.ErrCompanionAdsWrapperMissCompanions:
		return ResultCompanionAdsWrapperMissCompanions
	case vastbasic.ErrCreativeType:
		return ResultCreativeType
	case vastbasic.ErrCreativeWrapperType:
		return ResultCreativeWrapperType
	case vastbasic.ErrDurationEqualZero:
		return ResultDurationEqualZero
	case vastbasic.ErrDurationNegative:
		return ResultDurationNegative
	case vastbasic.ErrIconMissPosition:
		return ResultIconMissPosition
	case vastbasic.ErrIconMissProgram:
		return ResultIconMissProgram
	case vastbasic.ErrIconResourcesFormat:
		return ResultIconResourcesFormat
	case vastbasic.ErrImpressionMissURI:
		return ResultImpressionMissURI
	case vastbasic.ErrInlineMissAdTitle:
		return ResultInlineMissAdTitle
	case vastbasic.ErrInlineMissCreatives:
		return ResultInlineMissCreatives
	case vastbasic.ErrInlineMissImpressions:
		return ResultInlineMissImpressions
	case vastbasic.ErrLinearMissMediaFiles:
		return ResultLinearMissMediaFiles
	case vastbasic.ErrMediaFileHeightTooHigh:
		return ResultMediaFileHeightTooHigh
	case vastbasic.ErrMediaFileHeightTooLow:
		return ResultModiaFileHeightTooLow
	case vastbasic.ErrMediaFileMissDelivery:
		return ResultMediaFileMissDelivery
	case vastbasic.ErrMediaFileMissMimeType:
		return ResultMediaFileMissMimeType
	case vastbasic.ErrMediaFileMissURI:
		return ResultMediaFileMissURI
	case vastbasic.ErrMediaFileSize:
		return ResultMediaFileSize
	case vastbasic.ErrMediaFileUnsupportedMimeType:
		return ResultMediaFileUnsupportedMimeType
	case vastbasic.ErrMediaFileWidthTooHigh:
		return ResultMediaFileWidthTooHigh
	case vastbasic.ErrMediaFileWidthTooLow:
		return ResultMediaFileWidthTooLow
	case vastbasic.ErrNonLinearAdsMissNonLinears:
		return ResultNonLinearAdsMissNonLinears
	case vastbasic.ErrNonLinearResourceFormat:
		return ResultNonLinearResourceFormat
	case vastbasic.ErrOffsetPercentNegative:
		return ResultOffsetPercentNegative
	case vastbasic.ErrPricingCurrencyFormat:
		return ResultPricingCurrencyFormat
	case vastbasic.ErrPricingMissModel:
		return ResultPricingMissModel
	case vastbasic.ErrPricingMissPrice:
		return ResultPricingMissPrice
	case vastbasic.ErrUnsupportedDeliveryType:
		return ResultUnsupportedDeliveryType
	case vastbasic.ErrUnsupportedEvent:
		return ResultUnsupportedEvent
	case vastbasic.ErrUnsupportedMode:
		return ResultUnsupportedMode
	case vastbasic.ErrUnsupportedPriceModelType:
		return ResultUnsupportedPriceModelType
	case vastbasic.ErrUnsupportedVersion:
		return ResultUnsupportedVersion
	case vastbasic.ErrVastMissAd:
		return ResultVastMissAd
	case vastbasic.ErrVideoClicksMissClickThroughs:
		return ResultVideoClicksMissClickThroughs
	case vastbasic.ErrVideoDurationTooLong:
		return ResultVideoDurationTooLong
	case vastbasic.ErrVideoDurationTooShort:
		return ResultVideoDurationTooShort
	case vastbasic.ErrWrapperMissImpressions:
		return ResultWrapperMissImpressions
	case vastbasic.ErrWrapperMissVastAdTagURI:
		return ResultWrapperMissVastAdTagURI
	}

	return ResultUnknown
}

// EncodeInvalidVastHeader method encodes the invalid vast result into a string header value in the HTTP header.
func EncodeInvalidVastHeader(header http.Header, vastErr error) {
	validationResult := getValidationResultFromErr(vastErr)
	header.Set(HeaderXInvalidVastResult, strconv.Itoa(int(validationResult)))
}
