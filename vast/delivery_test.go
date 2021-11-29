package vast_test

import (
	"testing"

	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
)

func TestDeliveryValidateErrors(t *testing.T) {
	delivery1 := vast.Delivery("test")
	delivery2 := vast.Delivery(vast.DeliveryProgressive)
	delivery3 := vast.Delivery(vast.DeliveryStreaming)
	delivery4 := vast.Delivery("")
	vasttest.VerifyVastElementErrorAsExpected(t, delivery1, delivery1.Validate(), vast.ErrUnsupportedDeliveryType)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery2, delivery2.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery3, delivery3.Validate(), nil)
	vasttest.VerifyVastElementErrorAsExpected(t, delivery4, delivery4.Validate(), vast.ErrUnsupportedDeliveryType)
}
