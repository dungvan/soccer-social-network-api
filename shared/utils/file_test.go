package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistsFile(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	imageFilePath := dir + "/test/image/W_ULD_171201_01.jpg"
	fmt.Println(imageFilePath)
	isExsits := ExistsFile(imageFilePath)
	assert.True(t, isExsits)

	isExsits = ExistsFile("test")
	assert.False(t, isExsits)
}

func TestCopyFileOK(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	srcPath := dir + "/test/image/W_ULD_171201_01.jpg"
	dstPath := "/tmp/W_ULD_171201_01.jpg"

	err := CopyFile(srcPath, dstPath)
	assert.NoError(t, err)
	err = os.Remove(dstPath)
	assert.NoError(t, err)
}

func TestCopyFileCantOpenFile(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	srcPath := dir + "notfoundfile.txt"
	dstPath := "/tmp/W_ULD_171201_01.jpg"

	err := CopyFile(srcPath, dstPath)
	assert.Error(t, err)
}

func TestCopyFileCantDstFile(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	srcPath := dir + "/test/image/W_ULD_171201_01.jpga"
	dstPath := "/tmp/W_ULD_171201_01.jpg"

	f, err := os.OpenFile(dstPath, os.O_CREATE, 0444)
	assert.NoError(t, err)
	f.Close()

	err = CopyFile(srcPath, dstPath)
	assert.Error(t, err)

	err = os.Remove(dstPath)
	assert.NoError(t, err)
}

func TestReadByteFileOK(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	imagePath := dir + "/test/image/W_ULD_171201_01.jpg"

	b, err := ReadByteFile(imagePath)
	assert.NoError(t, err)
	assert.NotEmpty(t, b)
}

func TestReadByteFileNG(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	imagePath := dir + "notfoundfile.txt"

	b, err := ReadByteFile(imagePath)
	assert.Error(t, err)
	assert.Empty(t, b)
}

func TestFileDetectContentType(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"jp2 image file test",
			args{
				func() []byte {
					data := make([]byte, 128)
					file, _ := os.Open(dir + "/test/image/JP2.jp2")
					defer file.Close()
					file.Read(data)
					return data
				}(),
			},
		},
		{
			"jpx image file test",
			args{
				func() []byte {
					data := make([]byte, 128)
					file, _ := os.Open(dir + "/test/image/JPX.jpx")
					defer file.Close()
					file.Read(data)
					return data
				}(),
			},
		},
		{
			"pgm image file test",
			args{
				func() []byte {
					data := make([]byte, 128)
					file, _ := os.Open(dir + "/test/image/PGM.pgm")
					defer file.Close()
					file.Read(data)
					return data
				}(),
			},
		},
		{
			"pbm image file test",
			args{
				func() []byte {
					data := make([]byte, 128)
					file, _ := os.Open(dir + "/test/image/PBM.pbm")
					defer file.Close()
					file.Read(data)
					return data
				}(),
			},
		},
		{
			"tiff image file test",
			args{
				func() []byte {
					data := make([]byte, 128)
					file, _ := os.Open(dir + "/test/image/TIFF.tiff")
					defer file.Close()
					file.Read(data)
					return data
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectFileContentType(tt.args.data)
			assert.Contains(t, got, "image/", "want MIME type start by image/*, but got "+got)
		})
	}
}
