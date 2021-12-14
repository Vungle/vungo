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
	vasttest.VerifyVastElementErrorAsExpected(t, model1, model1.Validate(), vastbasic.ErrUnsupportedPriceModelType)
	vasttest.VerifyVastElementErrorAsExpected(t, model2, model2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model3, model3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model4, model4.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model5, model5.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model6, model6.Validate(), vastbasic.ErrUnsupportedPriceModelType)
}
