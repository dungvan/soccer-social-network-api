package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type gettag struct {
	a string `a:"1"`
}

func TestGetTagOK(t *testing.T) {
	testtag := &gettag{a: "test"}
	data, err := GetStructTag(testtag, "a")
	assert.Equal(t, "a:\"1\"", data)
	assert.NoError(t, err)
}

func TestGetTagNG(t *testing.T) {
	testtag := &gettag{a: "test"}
	data, err := GetStructTag(testtag, "b")
	assert.Equal(t, "", data)
	assert.Error(t, err)
}

func TestCreateHashFromPasswordOK(t *testing.T) {
	password := CreateHashFromPassword("test", "123456")
	assert.NotEmpty(t, password)
	assert.NotEqual(t, "123456", password)
}

func TestGetBrandPageURLFromImName(t *testing.T) {
	brandPageURLGU := GetBrandPageURLFromImName("2_image")
	assert.Equal(t, brandPrefixGU+"2_image", brandPageURLGU)
	brandPageURLGU = GetBrandPageURLFromImName("3_image")
	assert.Equal(t, brandPrefixGU+"3_image", brandPageURLGU)
	brandPageURLSTORE := GetBrandPageURLFromImName("image")
	assert.Equal(t, brandPrefixSTORE+"image", brandPageURLSTORE)
}

func TestGetWidthHeightImName(t *testing.T) {
	w, h := GetWidthHeightImName("test.jpg")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?r=1111&h=300")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=1111&r=300")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=a&h=300")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=1111&h=b")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=111111&h=300")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=1111&h=111111")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=0&h=111111")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=1111&h=0")
	assert.Equal(t, w, 0)
	assert.Equal(t, h, 0)
	w, h = GetWidthHeightImName("test.jpg?w=1111&h=300")
	assert.Equal(t, w, 1111)
	assert.Equal(t, h, 300)
}
