package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"testing"
)

func TestPricingModelValidateErrors(t *testing.T) {
	model1 := vast.PricingModel("cpp")
	model2 := vast.PricingModel(vast.PRICING_MODEL_CPC)
	model3 := vast.PricingModel(vast.PRICING_MODEL_CPE)
	model4 := vast.PricingModel(vast.PRICING_MODEL_CPM)
	model5 := vast.PricingModel(vast.PRICING_MODEL_CPV)
	model6 := vast.PricingModel("")
	vasttest.VerifyVastElementErrorAsExpected(t, model1, model1.Validate(), vast.ErrUnsupportedPriceModelType)
	vasttest.VerifyVastElementErrorAsExpected(t, model2, model2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model3, model3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model4, model4.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model5, model5.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model6, model6.Validate(), vast.ErrUnsupportedPriceModelType)
}
