package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vast3"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

func TestPricingModelValidateErrors(t *testing.T) {
	model1 := vast3.PricingModel("cpp")
	model2 := vast3.PricingModel(vast3.PricingModelCPC)
	model3 := vast3.PricingModel(vast3.PricingModelCPE)
	model4 := vast3.PricingModel(vast3.PricingModelCPM)
	model5 := vast3.PricingModel(vast3.PricingModelCPV)
	model6 := vast3.PricingModel("")
	vasttest.VerifyVastElementErrorAsExpected(t, model1, model1.Validate(), vastbasic.ErrUnsupportedPriceModelType)
	vasttest.VerifyVastElementErrorAsExpected(t, model2, model2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model3, model3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model4, model4.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model5, model5.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, model6, model6.Validate(), vastbasic.ErrUnsupportedPriceModelType)
}
