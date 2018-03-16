package validator

import (
	"reflect"

	v "gopkg.in/go-playground/validator.v9"
)

// tagFuncMaps: declare custom validation tag and function
var tagFuncMaps = map[string]func(v.FieldLevel) bool{
	"required_if":    RequiredIf,
	"s_eq":           SliceEq,
	"string_gt":      StringGt,
	"string_lt":      StringLt,
	"image_name":     isImageName,
	"hashtag":        isHashtag,
	"int_array_len":  intArrayLen,
	"max_array_len":  maxArrayLen,
	"source_im_name": sourceIMName,
}

var tagNameFunc = []func(reflect.StructField) string{
	FormTagName,
}

// New calls validator.New() and add custom validators
func New() *v.Validate {
	validate := v.New()
	for key, value := range tagFuncMaps {
		_ = validate.RegisterValidation(key, value)
	}
	for _, value := range tagNameFunc {
		validate.RegisterTagNameFunc(value)
	}
	return validate
}
