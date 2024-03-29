package vastelement_test

import (
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

var PricingModelType = reflect.TypeOf(vastelement.Pricing{})

func TestPricingMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Pricing", "pricing.xml", PricingModelType)
}

func TestPricingValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/pricing.xml", &vastelement.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm" currency="USD"></Pricing>`),
		&vastelement.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm">1.20</Pricing>`),
		&vastelement.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing currency="USD">1.20</Pricing>`),
		&vastelement.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="c" currency="USD">1.20</Pricing>`),
		&vastelement.Pricing{}, nil)
}
