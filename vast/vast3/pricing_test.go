package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast3"
	"reflect"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

var PricingModelType = reflect.TypeOf(vast3.Pricing{})

func TestPricingMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Pricing", "pricing.xml", PricingModelType)
}

func TestPricingValidateErrors(t *testing.T) {
	vasttest.VerifyVastElementFromFile(t, "testdata/pricing.xml", &vast3.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm" currency="USD"></Pricing>`),
		&vast3.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="cpm">1.20</Pricing>`),
		&vast3.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing currency="USD">1.20</Pricing>`),
		&vast3.Pricing{}, nil)
	vasttest.VerifyVastElementFromBytes(t, []byte(`<Pricing model="c" currency="USD">1.20</Pricing>`),
		&vast3.Pricing{}, nil)
}
