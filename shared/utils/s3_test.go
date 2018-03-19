package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetObjectPathOK(t *testing.T) {
	imageFile := "W_ULD_171201_01.jpg"
	dir := "test"
	path := GetObjectPath(StorageS3, dir, imageFile)
	assert.Equal(t, dir+"/"+imageFile, path)

	dir = "test/"
	path = GetObjectPath(StorageS3, dir, imageFile)
	assert.Equal(t, dir+imageFile, path)

	path = GetObjectPath(StorageMinio, dir, imageFile)
	assert.Equal(t, imageFile, path)
}

func TestGetStorageURLOKS3(t *testing.T) {
	imageFile := "W_ULD_171201_01.jpg"
	dir := "test"
	endpoint := "s3.amazonaws.com"
	bucketname := "dev-circle-photo"
	objectName := GetObjectPath(StorageS3, dir, imageFile)
	region := "ap-northeast-1"
	path := GetStorageURL(StorageS3, endpoint, true, bucketname, objectName, region)
	assert.Equal(t, "https://s3-ap-northeast-1.amazonaws.com/dev-circle-photo/test/W_ULD_171201_01.jpg", path)
}
func TestGetStorageURLOKMinio(t *testing.T) {
	imageFile := "W_ULD_171201_01.jpg"
	dir := "test"
	endpoint := "example.com"
	bucketname := "test"
	objectName := GetObjectPath(StorageMinio, dir, imageFile)
	region := ""
	path := GetStorageURL(StorageMinio, endpoint, true, bucketname, objectName, region)
	assert.Equal(t, "https://example.com/test/W_ULD_171201_01.jpg", path)

	path = GetStorageURL(StorageMinio, endpoint, false, bucketname, objectName, region)
	assert.Equal(t, "http://example.com/test/W_ULD_171201_01.jpg", path)
}
