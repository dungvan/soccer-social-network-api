package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	v "gopkg.in/go-playground/validator.v9"
)

type fieldLevelMock struct {
	top    reflect.Value
	parent reflect.Value
	field  reflect.Value
	param  reflect.Value
	str    string
}

func (f *fieldLevelMock) Top() reflect.Value {
	return f.top
}

// returns the current fields parent struct, if any or
// the comparison value if called 'VarWithValue'
func (f *fieldLevelMock) Parent() reflect.Value {
	return f.parent
}

// returns current field for validation
func (f *fieldLevelMock) Field() reflect.Value {
	return f.field
}

// returns the field's name with the tag
// name takeing precedence over the fields actual name.
func (f *fieldLevelMock) FieldName() string {
	return f.str
}

// returns the struct field's name
func (f *fieldLevelMock) StructFieldName() string {
	return f.str
}

// returns param for validation ag)ainst current field
func (f *fieldLevelMock) Param() string {
	return f.param.String()
}

func (f *fieldLevelMock) ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool) {
	return field, field.Kind(), false
}
func (f *fieldLevelMock) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	return f.param, f.param.Kind(), false
}

func TestRequiredIf(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	var nilSlice []byte
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if two params are not empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("value1"),
					param: reflect.ValueOf("value2"),
				},
			},
			true,
		},
		{
			"test return true if one of params is not empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf("value2"),
				},
			},
			true,
		},
		{
			"test return false if both the params are empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf(""),
				},
			},
			false,
		},
		{
			"test return true if one of params are empty and slice",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(nilSlice),
					param: reflect.ValueOf(""),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := RequiredIf(tt.args.fl); got != tt.want {
				t.Errorf("RequiredIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceEq(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if field contains param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf([]string{"a", "b"}),
					param: reflect.ValueOf("a@b"),
				},
			},
			true,
		},
		{
			"test return true if field contains param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf([]string{"a", "b"}),
					param: reflect.ValueOf("a@b@c"),
				},
			},
			true,
		},
		{
			"test return true if field not contains param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf([]string{"a", "d"}),
					param: reflect.ValueOf("a@b@c"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := SliceEq(tt.args.fl); got != tt.want {
				t.Errorf("SliceEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringGt(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if field > param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("1"),
					param: reflect.ValueOf("0"),
				},
			},
			true,
		},
		{
			"test return true if field = param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("1"),
					param: reflect.ValueOf("1"),
				},
			},
			false,
		},
		{
			"test return false if both the params are empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf("0"),
				},
			},
			false,
		},
		{
			"test return false if field is string",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("a"),
					param: reflect.ValueOf("0"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := StringGt(tt.args.fl); got != tt.want {
				t.Errorf("StringGt() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestStringLt(t *testing.T) {

	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		caseName string
		args     args
		want     bool
	}{
		{
			"test return true if field < param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("-1"),
					param: reflect.ValueOf("0"),
				},
			},
			true,
		},
		{
			"test return true if field = param",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("1"),
					param: reflect.ValueOf("1"),
				},
			},
			false,
		},
		{
			"test return false if both the params are empty",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf(""),
					param: reflect.ValueOf("0"),
				},
			},
			false,
		},
		{
			"test return false if field is string",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("a"),
					param: reflect.ValueOf("0"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := StringLt(tt.args.fl); got != tt.want {
				t.Errorf("StringGt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsImageName(t *testing.T) {
	type args struct {
		fl v.FieldLevel
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test return true if ImageName has imagename value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("imagename"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has imagename123 value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("imagename123"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has image-name.jpg value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image-name.jpg"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has image-_2name.jpg value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image-_2name.jpg"),
				},
			},
			true,
		},
		{
			"test return true if ImageName has 403540-69 value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("403540-69"),
				},
			},
			true,
		},
		{
			"test return false if ImageName has image# value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image#"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image@ value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image@"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image! value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image!"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image% value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image%"),
				},
			},
			false,
		},
		{
			"test return false if ImageName has image&* value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("image&*"),
				},
			},
			false,
		},
		{
			"test return true if false has @$123 value",
			args{
				&fieldLevelMock{
					field: reflect.ValueOf("@$123"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := isImageName(tt.args.fl)
			{
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestIsHashtagReturnTrue(t *testing.T) {
	type args struct {
		fl v.FieldLevel
	}
	caseTrue := []string{
		"０９",
		"hashtag",
		"hashtag1",
		"1hashtag",
		"hashtag_",
		"_hashtag",
		"1_hashtag",
		"hashtag_1",
		"123123",
		"ぁあぃいぅ",  //Hiragana
		"ァアィイゥウ", //Katakana (Full Width)
		"漢字日本語",  //Kanji
		"hashtag_123_ぁあ_ぁあぃ_ァア_漢字日_０９",
		"ㇰㇱㇲ",   //Miscellaneous Japanese Symbols and Characters
		"ｦ",     //Katakana and Punctuation (Half Width)
		"ｦｧｨｩｪ", //Katakana and Punctuation (Half Width)
		"ㇿ",
		"0000",
	}

	for _, v := range caseTrue {
		t.Run("test return TRUE with value "+v, func(t *testing.T) {
			args := args{&fieldLevelMock{
				field: reflect.ValueOf(v),
			},
			}
			got := isHashtag(args.fl)
			{
				assert.True(t, got)
			}
		})
	}
}

func TestIsHashtagReturnFalse(t *testing.T) {
	type args struct {
		fl v.FieldLevel
	}

	caseFalse := []string{
		" ",
		"#hashtag",
		"hashtag#",
		"h#ashtag",
		"hashta#g",
		" #hashtag",
		"#hash=tag1",
		" hashtag",
		"h ashtag",
		"hashtag ",
		"hash@tag1",
		"hash,tag1",
		"hashtag 1",
		"1 hashtag",
		"-hashtag",
		"hash-tag",
		"!@#$%^&*",
		"【】〒〓〔〕", //Japanese Symbols and Punctuation
		"㊇㊈㊉㊊",   //Miscellaneous Japanese Symbols and Characters
		"・",      //Katakana (Full Width)
		"゠ァアィイゥウ", //Katakana (Full Width),
		"｟", //Katakana and Punctuation (Half Width)
		"｟｠｡｢｣､･",   //Katakana and Punctuation (Half Width)
		"⾍⾎⾏⺀⺁⺂⺃⺄⺅", //Kanji Radicals,
		" ゚",
		"゛",
		"゜",
		"⺀",
		"⿕",
		"ﾞ",
		"、",
		"〾",
		"㈠",
		"㍿",
		":D",
		"🤖",
		"🇩🇪", // 8byte
	}

	for _, v := range caseFalse {
		t.Run("test return FALSE with value "+v, func(t *testing.T) {

			args := args{&fieldLevelMock{
				field: reflect.ValueOf(v),
			},
			}
			got := isHashtag(args.fl)
			{
				assert.False(t, got)
			}
		})
	}
}
func TestSourceIMNameReturnFalse(t *testing.T) {
	type args struct {
		fl v.FieldLevel
	}

	caseFalse := []string{
		" ",
		"image,.jpg",
		"image-.jpg",
		"🤖",
		"シャツ.jpg",
	}

	for _, v := range caseFalse {
		t.Run("test return FALSE with value "+v, func(t *testing.T) {

			args := args{&fieldLevelMock{
				field: reflect.ValueOf(v),
			},
			}
			got := sourceIMName(args.fl)
			{
				assert.False(t, got)
			}
		})
	}
}

func TestIntArrayLenFunc(t *testing.T) {
	testCases := []struct {
		caseName string
		args     *fieldLevelMock
		expected bool
	}{
		{
			"test return TRUE when array len valid",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{1, 2, 2, 4}),
				param: reflect.ValueOf("4"),
			},
			true,
		},
		{
			"test return FALSE when array len > param",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{1, 2, 2, 4, 5}),
				param: reflect.ValueOf("4"),
			},
			false,
		},
		{
			"test return FALSE when array len < param",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{1, 2, 2}),
				param: reflect.ValueOf("4"),
			},
			false,
		},
		{
			"test return FALSE when when this type is string",
			&fieldLevelMock{
				field: reflect.ValueOf([]string{"1", "2", "3", "4"}),
				param: reflect.ValueOf("4"),
			},
			false,
		},
		{
			"test return FALSE when when this type is array of float",
			&fieldLevelMock{
				field: reflect.ValueOf([]float64{1, 2, 3, 4}),
				param: reflect.ValueOf("4"),
			},
			false,
		},
		{
			"test return FALSE when input is not an array",
			&fieldLevelMock{
				field: reflect.ValueOf("string"),
				param: reflect.ValueOf("4"),
			},
			false,
		},
		{
			"test return FALSE when params invalid",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{11}),
				param: reflect.ValueOf("4x"),
			},
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			actual := intArrayLen(testCase.args)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestSourceIMNameReturnTrue(t *testing.T) {
	type args struct {
		fl v.FieldLevel
	}

	caseTrue := []string{
		"image",
		"image...___",
		"image.jpg",
		"image1.jpg",
		"image_.jpg",
		"_image.jpg",
		"image_1.jpg",
	}

	for _, v := range caseTrue {
		t.Run("test return TRUE with value "+v, func(t *testing.T) {

			args := args{&fieldLevelMock{
				field: reflect.ValueOf(v),
			},
			}
			got := sourceIMName(args.fl)
			{
				assert.True(t, got)
			}
		})
	}
}

func TestMaxArrayLenFunc(t *testing.T) {
	testCases := []struct {
		caseName string
		args     *fieldLevelMock
		expected bool
	}{
		{
			"test return TRUE when array len < params",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{1, 2, 2}),
				param: reflect.ValueOf("4"),
			},
			true,
		},
		{
			"test return TRUE when array len = params",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{1, 2, 2, 4}),
				param: reflect.ValueOf("4"),
			},
			true,
		},
		{
			"test return FALSE when array len > params",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{1, 2, 2, 2, 2}),
				param: reflect.ValueOf("4"),
			},
			false,
		},
		{
			"test return FALSE when input invalid",
			&fieldLevelMock{
				field: reflect.ValueOf("is strings"),
				param: reflect.ValueOf("4"),
			},
			false,
		},
		{
			"test return FALSE when param invalid",
			&fieldLevelMock{
				field: reflect.ValueOf([]uint{1, 2, 2, 2, 2}),
				param: reflect.ValueOf("4x"),
			},
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			actual := maxArrayLen(testCase.args)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
