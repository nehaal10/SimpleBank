package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nehaal10/simeplebank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsdupportedCurrency(currency)
	}
	return false
}
