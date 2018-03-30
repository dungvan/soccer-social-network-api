package validator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	v "gopkg.in/go-playground/validator.v9"
)

const (
	numberRegexString = "^-?[0-9]+$"
)

var (
	numberRegex = regexp.MustCompile(numberRegexString)
)

// RequiredIf implements validator.Func
func RequiredIf(fl v.FieldLevel) bool {
	currentValue, currentKind, _ := fl.ExtractType(fl.Field())
	reflValue, reflKind, _ := fl.GetStructFieldOK()
	// The both fields are available
	if HasValue(currentValue, currentValue.Type(), currentKind) {
		return true
	}

	// One of the fields is available
	if HasValue(reflValue, reflValue.Type(), reflKind) {
		return true
	}
	// The both fields are not available
	return false
}

// SliceEq performs Slice type match check.
// Slice type only validation. Other type are always false.
func SliceEq(fl v.FieldLevel) bool {
	_, currentKind, _ := fl.ExtractType(fl.Field())
	field := fl.Field()
	param := fl.Param()

	// split @("or separate")
	params := strings.Split(param, "@")

	var isMatch bool

	switch currentKind {
	case reflect.Slice:
		data := field.Interface().([]string)
		for _, v := range data {
			isMatch = false
			for _, p := range params {
				if v == p {
					isMatch = true
					break
				}
			}
			if !isMatch {
				return false
			}
		}
		return true
	}
	return false
}

// StringGt change string to int and ">" check.
func StringGt(fl v.FieldLevel) bool {
	_, currentKind, _ := fl.ExtractType(fl.Field())
	param := fl.Param()

	switch currentKind {
	case reflect.String:
		value, err := asInt(fl.Field().String())
		if err != nil {
			return false
		}
		p, err := asInt(param)
		if err != nil {
			return false
		}
		return value > p
	}
	return false
}

// StringLt change string to int and "<" check.
func StringLt(fl v.FieldLevel) bool {
	_, currentKind, _ := fl.ExtractType(fl.Field())
	param := fl.Param()

	switch currentKind {
	case reflect.String:
		value, err := asInt(fl.Field().String())
		if err != nil {
			return false
		}
		p, err := asInt(param)
		if err != nil {
			return false
		}
		return value < p
	}
	return false
}

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 64)
	return i, err
}

// isImageName is the validation function for validating if the current field's value is a valid imagename value.
func isImageName(fl v.FieldLevel) bool {
	return imageNameRegex.MatchString(fl.Field().String())
}

// isHashtag is the validation function for validating if a value is a valid hashtag.
func isHashtag(fl v.FieldLevel) bool {
	return hashtagRegex.MatchString(fl.Field().String())
}

// arrayNumLen is the validation function for validating if the current field's value is a valid length value.
func intArrayLen(fl v.FieldLevel) bool {
	_, currentKind, _ := fl.ExtractType(fl.Field())
	param := fl.Param()
	switch currentKind {
	case reflect.Slice:
		p, err := asInt(param)
		if err != nil {
			return false
		}
		length := fl.Field().Len()
		x := int64(length)
		if x != p {
			return false
		}
		for i := 0; i < fl.Field().Len(); i++ {
			_, eType, _ := fl.ExtractType(fl.Field().Index(i))
			if eType != reflect.Uint {
				return false
			}
		}
		return true
	}
	return false
}

// maxArrayLen is the validation function for validating len of current field's value.
func maxArrayLen(fl v.FieldLevel) bool {
	_, currentKind, _ := fl.ExtractType(fl.Field())
	param := fl.Param()
	switch currentKind {
	case reflect.Slice:
		p, err := asInt(param)
		if err != nil {
			return false
		}
		if int64(fl.Field().Len()) <= p {
			return true
		}
	}
	return false
}

func sourceIMName(fl v.FieldLevel) bool {
	return sourceIMNameRegex.MatchString(fl.Field().String())
}

func sourceVideoName(fl v.FieldLevel) bool {
	return sourceVideoNameRegex.MatchString(fl.Field().String())
}
