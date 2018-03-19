package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormTagName(t *testing.T) {
	type args struct {
		fld reflect.StructField
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"struct with form tag",
			args{
				reflect.StructField{
					Tag:  reflect.StructTag(`form:"test_name"`),
					Name: "TestName",
				},
			},
			"test_name",
		},
		{
			"struct with no form & json tag",
			args{
				reflect.StructField{
					Name: "TestName",
				},
			},
			"TestName",
		},
		{
			"struct with json tag",
			args{
				reflect.StructField{
					Tag:  reflect.StructTag(`json:"test_name"`),
					Name: "TestName",
				},
			},
			"test_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormTagName(tt.args.fld)
			assert.Equal(t, tt.want, got)
		})
	}
}
