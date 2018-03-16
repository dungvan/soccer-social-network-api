package storage

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Request interface
type Request interface {
	UploadToS3() (*s3.PutObjectOutput, error)
	SetParam(imageFile multipart.File, bucketName string, objectName string, fileType string, acl ...ACL) Request
	IsSetParam() bool
}

// ACL type
type ACL string

var (
	// BucketCannedACLPublicReadWrite public read write
	BucketCannedACLPublicReadWrite = ACL(s3.BucketCannedACLPublicReadWrite)
	// BucketCannedACLPrivate private
	BucketCannedACLPrivate = ACL(s3.BucketCannedACLPrivate)
	// BucketCannedACLPublicRead public read
	BucketCannedACLPublicRead = ACL(s3.BucketCannedACLPublicRead)
	// BucketCannedACLAuthenticatedRead authenticated read
	BucketCannedACLAuthenticatedRead = ACL(s3.BucketCannedACLAuthenticatedRead)
)

// uploadS3Request struct
type uploadS3Request struct {
	s3     *s3.S3
	params *params
}

type params struct {
	ImageFile  multipart.File
	BucketName string
	ObjectName string
	FileType   string
	ACL        ACL
}

// UploadToS3 function upload image into s3
func (r *uploadS3Request) UploadToS3() (*s3.PutObjectOutput, error) {
	putObj, err := r.s3.PutObject(&s3.PutObjectInput{
		Body:        r.params.ImageFile,
		Bucket:      aws.String(r.params.BucketName),
		Key:         aws.String(r.params.ObjectName),
		ContentType: aws.String(r.params.FileType),
		ACL:         aws.String(string(r.params.ACL)),
	})
	r.params = nil
	return putObj, err
}

// SetParam set param for uploadS3Request
func (r *uploadS3Request) SetParam(imageFile multipart.File, bucketName string, objectName string, fileType string, acl ...ACL) Request {
	if !r.IsSetParam() {
		r.params = &params{}
	}
	r.params.ImageFile = imageFile
	r.params.BucketName = bucketName
	r.params.ObjectName = objectName
	r.params.FileType = fileType
	if len(acl) > 0 {
		r.params.ACL = acl[0]
	}
	return r
}

// IsSetParam upload S3 request
func (r *uploadS3Request) IsSetParam() bool {
	if r.params != nil {
		return true
	}
	return false
}

// New a UploadS3Request struct
func New(s3Client *s3.S3) Request {
	r := &uploadS3Request{s3: s3Client}
	return r
}
