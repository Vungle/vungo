package vast2_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var PricingModelType = reflect.TypeOf(vast2.Pricing{})

func TestPricingMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Pricing", "pricing.xml", PricingModelType)
}

func TestPricingValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/pricing.xml", &vast2.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm" currency="USD"></Pricing>`),
		&vast2.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm">1.20</Pricing>`),
		&vast2.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing currency="USD">1.20</Pricing>`),
		&vast2.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="c" currency="USD">1.20</Pricing>`),
		&vast2.Pricing{}, nil)
}
