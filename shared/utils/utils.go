package utils

import (
	"encoding/hex"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/crypto/scrypt"
)

var (
	brandPrefixGU    = "www.uniqlo.com/jp/gu/item/"
	brandPrefixSTORE = "www.uniqlo.com/jp/store/goods/"
)

// GetStructTag get struct tag.
func GetStructTag(typeData interface{}, tagName string) (string, error) {
	field, ok := reflect.TypeOf(typeData).Elem().FieldByName(tagName)
	if !ok {
		return "", ErrorsNew("can't get Tag")
	}
	return string(field.Tag), nil
}

// CreateHashFromPassword scrypt hash password.
func CreateHashFromPassword(salt, password string) string {
	converted, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:])
}

// GetBrandPageURLFromImName return BrandPageURLfrom imName.
func GetBrandPageURLFromImName(imName string) string {
	if strings.HasPrefix(imName, "2") || strings.HasPrefix(imName, "3") {
		return brandPrefixGU + imName
	}
	return brandPrefixSTORE + imName
}

// GetWidthHeightImName return width and height of image from sourceImageFileName.
func GetWidthHeightImName(sourceImageFileName string) (width int, height int) {
	strSplit := strings.FieldsFunc(sourceImageFileName, func(c rune) bool { return c == '=' || c == '&' || c == '?' })
	n := len(strSplit)
	if n < 5 || strSplit[n-2] != "h" || strSplit[n-4] != "w" || len(strSplit[n-1]) > 5 || len(strSplit[n-3]) > 5 {
		return
	}
	height, err := strconv.Atoi(strSplit[n-1])
	if height == 0 || err != nil {
		return
	}
	width, err = strconv.Atoi(strSplit[n-3])
	if width == 0 || err != nil {
		return 0, 0
	}
	return
}
