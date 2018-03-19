package utils

import (
	"regexp"
	"strings"
)

const (
	// StorageS3 is s3.
	StorageS3 = "s3"
	// StorageMinio is minio.
	StorageMinio = "minio"
)

// GetObjectPath get objectpath.
func GetObjectPath(storage, s3Path, uploadFile string) string {
	if storage == StorageS3 {
		re := regexp.MustCompile(`/$`)
		var addSlash string
		if !re.MatchString(s3Path) {
			addSlash = "/"
		}
		return s3Path + addSlash + uploadFile
	}
	return uploadFile
}

// GetStorageURL get s3 or minio url.
func GetStorageURL(storage string, endpoint string, secure bool, bucketName string, objectName string, region string) string {
	if storage == StorageS3 {
		return getS3URL(endpoint, secure, bucketName, objectName, region)
	}
	return getMinioURL(endpoint, secure, bucketName, objectName)
}

// getS3URL returns aws s3 url.
func getS3URL(endpoint string, secure bool, bucketName string, objectName string, region string) string {
	protocol := getProtocol(secure)
	return strings.Replace(endpoint, "s3", protocol+"s3-"+region, 1) + "/" + bucketName + "/" + objectName
}

// getMinioURL returns minio url.
func getMinioURL(endpoint string, secure bool, bucketName string, objectName string) string {
	protocol := getProtocol(secure)
	return protocol + endpoint + "/" + bucketName + "/" + objectName
}

func getProtocol(secure bool) string {
	if secure {
		return "https://"
	}
	return "http://"
}
