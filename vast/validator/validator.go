package validator

import "github.com/Vungle/vungo/vast/entity"

type Validator interface {
	ValidateVast(v *entity.Vast) error
}
