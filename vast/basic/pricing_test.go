package vastbasic_test

import (
	"reflect"
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"

	"github.com/Vungle/vungo/vast/vasttest"
)

var PricingModelType = reflect.TypeOf(vastbasic.Pricing{})

func TestPricingMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Pricing", "pricing.xml", PricingModelType)
}

func TestPricingValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/pricing.xml", &vastbasic.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm" currency="USD"></Pricing>`),
		&vastbasic.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm">1.20</Pricing>`),
		&vastbasic.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing currency="USD">1.20</Pricing>`),
		&vastbasic.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="c" currency="USD">1.20</Pricing>`),
		&vastbasic.Pricing{}, nil)
}
