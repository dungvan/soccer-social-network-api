package storage

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/stretchr/testify/assert"
)

var (
	s3Config = &aws.Config{
		Region:           aws.String("ap-southeast-1"),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials("AKIAJAI2S6FXF5ERTXUA", "uNy5xvHoqVSFHaL9mcDP75nnQ+Z+UIP9rG8iCwM/", ""),
	}
	sess       = session.Must(session.NewSession(s3Config))
	s3Client   = s3.New(sess)
	dstFile    = "/tmp/text.txt"
	fileName   = "text.txt"
	s3Path     = "item"
	fileType   = "text/plain; charset=utf-8"
	objectName = utils.GetObjectPath("s3", s3Path, fileName)
	bucketName = "yoshiaki-miyazaki-test"
)

func TestUploadToS3Success(t *testing.T) {
	file, err := os.OpenFile(dstFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		t.Errorf("Can read file %s", err)
	}
	r := &uploadS3Request{
		s3: s3Client,
		params: &params{
			ACL:        BucketCannedACLPublicReadWrite,
			FileType:   fileType,
			ImageFile:  file,
			ObjectName: objectName,
			BucketName: bucketName,
		},
	}
	putObj, err := r.UploadToS3()
	assert.NoError(t, err)
	assert.NotEmpty(t, putObj)
}

func TestUploadToS3Fail(t *testing.T) {
	file, err := os.OpenFile(dstFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		t.Errorf("Can read file %s", err)
	}
	tests := []struct {
		description     string
		uploadS3Request *uploadS3Request
	}{
		{
			"missing bucket name",
			&uploadS3Request{
				s3: s3Client,
				params: &params{
					ACL:        BucketCannedACLPublicReadWrite,
					FileType:   fileType,
					ImageFile:  file,
					ObjectName: objectName,
				},
			},
		},
		{
			"wrong bucket name",
			&uploadS3Request{
				s3: s3Client,
				params: &params{
					ACL:        BucketCannedACLPublicReadWrite,
					FileType:   fileType,
					ImageFile:  file,
					ObjectName: objectName,
					BucketName: "wrong_bucketName",
				},
			},
		},
		{
			"missing object name",
			&uploadS3Request{
				s3: s3Client,
				params: &params{
					ACL:        BucketCannedACLPublicReadWrite,
					FileType:   fileType,
					ImageFile:  file,
					BucketName: bucketName,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			putObj, err := tt.uploadS3Request.UploadToS3()
			assert.Error(t, err)
			assert.Equal(t, &s3.PutObjectOutput{}, putObj)
		})
	}
}

func TestSetParam(t *testing.T) {
	file, err := os.OpenFile(dstFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		t.Errorf("Can read file %s", err)
	}
	tests := []struct {
		description     string
		uploadS3Request *uploadS3Request
		acl             ACL
	}{
		{
			"IsSetParam = true",
			&uploadS3Request{
				s3:     s3Client,
				params: &params{},
			},
			"",
		},
		{
			"IsSetParam = false",
			&uploadS3Request{
				s3: s3Client,
			},
			"",
		},
		{
			"len(acl) > 0",
			&uploadS3Request{
				s3: s3Client,
			},
			BucketCannedACLPublicReadWrite,
		},
		{
			"len(acl) < 0",
			&uploadS3Request{
				s3: s3Client,
			},
			"",
		},
		{
			"len(acl) = 0",
			&uploadS3Request{
				s3: s3Client,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.uploadS3Request.SetParam(file, bucketName, objectName, fileType, tt.acl)
			assert.NotNil(t, got)
		})
	}
}

func TestNew(t *testing.T) {
	got := New(s3Client)
	assert.NotNil(t, got)
}

func TestIsSetParam(t *testing.T) {
	tests := []struct {
		description     string
		uploadS3Request *uploadS3Request
		isSetParam      bool
	}{
		{
			"params != nil",
			&uploadS3Request{
				s3:     s3Client,
				params: &params{},
			},
			true,
		},
		{
			"params = nil",
			&uploadS3Request{
				s3: s3Client,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.uploadS3Request.IsSetParam()
			assert.NotNil(t, got)
		})
	}
}
