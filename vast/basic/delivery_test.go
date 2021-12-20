package vastbasic_test

import (
	"testing"

	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestDeliveryValidateErrors(t *testing.T) {
	delivery1 := vastbasic.Delivery("test")
	delivery2 := vastbasic.Delivery(vastbasic.DeliveryProgressive)
	delivery3 := vastbasic.Delivery(vastbasic.DeliveryStreaming)
	delivery4 := vastbasic.Delivery("")
	vasttest.VerifyVastElementErrorAsExpected(t, delivery1, vasttest.ValidateElement(delivery1), vastbasic.ErrUnsupportedDeliveryType)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery2, vasttest.ValidateElement(delivery2), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery3, vasttest.ValidateElement(delivery3), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery4, vasttest.ValidateElement(delivery4), vastbasic.ErrUnsupportedDeliveryType)
}
