package vast3_test

import (
	"github.com/Vungle/vungo/vast/vast2"
	"testing"

	"github.com/Vungle/vungo/vast/vasttest"
)

func TestDeliveryValidateErrors(t *testing.T) {
	delivery1 := vast2.Delivery("test")
	delivery2 := vast2.Delivery(vast2.DeliveryProgressive)
	delivery3 := vast2.Delivery(vast2.DeliveryStreaming)
	delivery4 := vast2.Delivery("")
	vasttest.VerifyVastElementErrorAsExpected(t, delivery1, delivery1.Validate(), vast2.ErrUnsupportedDeliveryType)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery2, delivery2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery3, delivery3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery4, delivery4.Validate(), vast2.ErrUnsupportedDeliveryType)
}
