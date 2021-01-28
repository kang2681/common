package validate

import (
	"github.com/asaskevich/govalidator"
)

func IsEmail(str string) bool {
	return govalidator.IsEmail(str)
}
