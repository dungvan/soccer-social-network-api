package validator

import (
	"reflect"
)

// FormTagName for register new name field by tag
func FormTagName(fld reflect.StructField) string {
	name := fld.Tag.Get("form")
	if name != "" {
		return name
	}
	name = fld.Tag.Get("json")
	if name != "" {
		return name
	}
	return fld.Name
}
