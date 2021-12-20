package vastbasic_test

import (
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestPricingModelValidateErrors(t *testing.T) {
	model1 := vastbasic.PricingModel("cpp")
	model2 := vastbasic.PricingModel(vastbasic.PricingModelCPC)
	model3 := vastbasic.PricingModel(vastbasic.PricingModelCPE)
	model4 := vastbasic.PricingModel(vastbasic.PricingModelCPM)
	model5 := vastbasic.PricingModel(vastbasic.PricingModelCPV)
	model6 := vastbasic.PricingModel("")
	vasttest.VerifyVastElementErrorAsExpected(t, model1, vasttest.ValidateElement(model1), vastbasic.ErrUnsupportedPriceModelType)
	vasttest.VerifyVastElementErrorAsExpected(t, model2, vasttest.ValidateElement(model2), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model3, vasttest.ValidateElement(model3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model4, vasttest.ValidateElement(model4), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model5, vasttest.ValidateElement(model5), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model6, vasttest.ValidateElement(model6), vastbasic.ErrUnsupportedPriceModelType)
}
