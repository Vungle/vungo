package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestPricingModelValidateErrors(t *testing.T) {
	model1 := vastelement.PricingModel("cpp")
	model2 := vastelement.PricingModel(vastelement.PricingModelCPC)
	model3 := vastelement.PricingModel(vastelement.PricingModelCPE)
	model4 := vastelement.PricingModel(vastelement.PricingModelCPM)
	model5 := vastelement.PricingModel(vastelement.PricingModelCPV)
	model6 := vastelement.PricingModel("")
	vasttest.VerifyVastElementErrorAsExpected(t, model1, model1.Validate(vastelement.Version3), vastelement.ErrUnsupportedPriceModelType)
	vasttest.VerifyVastElementErrorAsExpected(t, model2, model2.Validate(vastelement.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model3, model3.Validate(vastelement.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model4, model4.Validate(vastelement.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model5, model5.Validate(vastelement.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model6, model6.Validate(vastelement.Version3), vastelement.ErrUnsupportedPriceModelType)
}
