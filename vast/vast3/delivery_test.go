package vast3_test

import (
	"github.com/Vungle/vungo/vast/basic"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

func TestDeliveryValidateErrors(t *testing.T) {
	delivery1 := vastbasic.Delivery("test")
	delivery2 := vastbasic.Delivery(vastbasic.DeliveryProgressive)
	delivery3 := vastbasic.Delivery(vastbasic.DeliveryStreaming)
	delivery4 := vastbasic.Delivery("")
	vasttest.VerifyVastElementErrorAsExpected(t, delivery1, delivery1.Validate(), vastbasic.ErrUnsupportedDeliveryType)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery2, delivery2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery3, delivery3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery4, delivery4.Validate(), vastbasic.ErrUnsupportedDeliveryType)
}
