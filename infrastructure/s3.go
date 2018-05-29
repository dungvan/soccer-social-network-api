package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dungvan/soccer-social-network-api/shared/storage"
)

// S3 struct
type S3 struct {
	*s3.S3
}

// NewS3RequestFunc func
type NewS3RequestFunc func() storage.Request

var (
	// Storage storage s3
	Storage string
	// Endpoint endpoint s3
	Endpoint string
	// Secure secure s3
	Secure bool
	// Region region s3
	Region string
)

func init() {
	Storage = GetConfigString("objectstorage.storage")
	Endpoint = GetConfigString("objectstorage.endpoint")
	Secure = GetConfigBool("objectstorage.secure")
	Region = GetConfigString("objectstorage.region")
}

const (
	// StorageS3 is s3.
	StorageS3 = "s3"
	// StorageMinio is minio.
	StorageMinio = "minio"
)

// NewS3 returns new S3 instance.
func NewS3() *S3 {
	accessKeyID := GetConfigString("objectstorage.accesskey")
	secretAccessKey := GetConfigString("objectstorage.secretkey")

	var s3Config *aws.Config
	switch Storage {
	case StorageS3:
		// Configure to use Server
		s3Config = &aws.Config{
			Region:           aws.String(Region),
			DisableSSL:       aws.Bool(!Secure),
			S3ForcePathStyle: aws.Bool(true),
		}
		break
	case StorageMinio:
		// Configure to use Server
		s3Config = &aws.Config{
			Endpoint:         aws.String(Endpoint),
			Region:           aws.String(Region),
			DisableSSL:       aws.Bool(!Secure),
			S3ForcePathStyle: aws.Bool(true),
		}
		break
	default:
		panic("storage must select " + StorageS3 + " or " + StorageMinio)
	}

	s3Config.Credentials = credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	sess := session.Must(session.NewSession(s3Config))
	s3Client := s3.New(sess)

	return &S3{s3Client}
}

// NewRequest upload s3
func (s3 *S3) NewRequest() storage.Request {
	return storage.New(s3.S3)
}
