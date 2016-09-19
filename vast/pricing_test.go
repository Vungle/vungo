package vast_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

var PricingModelType = reflect.TypeOf(vast.Pricing{})

func TestPricingMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Pricing", "pricing.xml", PricingModelType)
}

func TestPricingValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/pricing.xml", &vast.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm" currency="USD"></Pricing>`),
		&vast.Pricing{}, vast.ErrPricingMissPrice)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm">1.20</Pricing>`),
		&vast.Pricing{}, vast.ErrPricingCurrencyFormat)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing currency="USD">1.20</Pricing>`),
		&vast.Pricing{}, vast.ErrPricingMissModel)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="c" currency="USD">1.20</Pricing>`),
		&vast.Pricing{}, vast.ErrUnsupportedPriceModelType)
}
