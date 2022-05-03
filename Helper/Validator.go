package Helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var Valid *validator.Validate = validator.New()

func ValidStruct(s interface{}) []error {
	errMsg := make([]error, 0)
	err := Valid.Struct(s)
	if err != nil {
		fmt.Println(err)
		validErr := err.(validator.ValidationErrors)
		for _, e := range validErr {
			err = errors.New(fmt.Sprintf("%s must be %s", e.Field(), e.Tag()))

			errMsg = append(errMsg, err)

			return errMsg

		}

	}
	return nil

}
