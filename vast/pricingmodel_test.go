package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestPricingModelValidateErrors(t *testing.T) {
	model1 := vast.PricingModel("cpp")
	model2 := vast.PricingModel(vast.PricingModelCPC)
	model3 := vast.PricingModel(vast.PricingModelCPE)
	model4 := vast.PricingModel(vast.PricingModelCPM)
	model5 := vast.PricingModel(vast.PricingModelCPV)
	model6 := vast.PricingModel("")
	vasttest.VerifyVastElementErrorAsExpected(t, model1, model1.Validate(), vast.ErrUnsupportedPriceModelType)
	vasttest.VerifyVastElementErrorAsExpected(t, model2, model2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model3, model3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model4, model4.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model5, model5.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model6, model6.Validate(), vast.ErrUnsupportedPriceModelType)
}
