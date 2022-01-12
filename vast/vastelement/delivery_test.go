package vastelement_test

import (
	"testing"

	"github.com/Vungle/vungo/vast/vastelement"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestDeliveryValidateErrors(t *testing.T) {
	delivery1 := vastelement.Delivery("test")
	delivery2 := vastelement.Delivery(vastelement.DeliveryProgressive)
	delivery3 := vastelement.Delivery(vastelement.DeliveryStreaming)
	delivery4 := vastelement.Delivery("")
	vasttest.VerifyVastElementErrorAsExpected(t, delivery1, delivery1.Validate(vastelement.Version3), vastelement.ErrUnsupportedDeliveryType)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery2, delivery2.Validate(vastelement.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery3, delivery3.Validate(vastelement.Version3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery4, delivery4.Validate(vastelement.Version3), vastelement.ErrUnsupportedDeliveryType)
}
