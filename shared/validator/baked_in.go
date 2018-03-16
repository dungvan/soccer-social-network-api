package validator

import "reflect"

// HasValue is the validation function for validating if the current field's value is not the default static value.
// NOTE: This is exposed for use within your own custom functions and not intended to be called directly.
func HasValue(field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind) bool {
	switch fieldKind {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(fieldType).Interface()
	}
}
