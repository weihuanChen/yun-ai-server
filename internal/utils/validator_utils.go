package utils

import validator2 "github.com/go-playground/validator/v10"

type (
	ValidErrRes struct {
		Error bool
		Field string
		Tag   string
		Value interface{}
	}
)

var validator = validator2.New()

// 参数验证器
func Validator(data interface{}) []ValidErrRes {
	var Errors []ValidErrRes
	errs := validator.Struct(data)
	if errs != nil {
		// 类型断言并获取
		for _, err := range errs.(validator2.ValidationErrors) {
			var el ValidErrRes
			el.Error = true
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Value()

			Errors = append(Errors, el)
		}
	}
	return Errors
}
