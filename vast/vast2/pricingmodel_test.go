package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

func TestPricingModelValidateErrors(t *testing.T) {
	model1 := vast2.PricingModel("cpp")
	model2 := vast2.PricingModel(vast2.PricingModelCPC)
	model3 := vast2.PricingModel(vast2.PricingModelCPE)
	model4 := vast2.PricingModel(vast2.PricingModelCPM)
	model5 := vast2.PricingModel(vast2.PricingModelCPV)
	model6 := vast2.PricingModel("")
	vasttest.VerifyVastElementErrorAsExpected(t, model1, model1.Validate(), vast2.ErrUnsupportedPriceModelType)
	vasttest.VerifyVastElementErrorAsExpected(t, model2, model2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model3, model3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model4, model4.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model5, model5.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model6, model6.Validate(), vast2.ErrUnsupportedPriceModelType)
}
