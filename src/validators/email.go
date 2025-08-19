package validators

import (
	"github.com/asaskevich/govalidator"
)

func IsEmail(s string) bool {
	return govalidator.IsEmail(s)
}
