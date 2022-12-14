package fileupload

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
}

func NewS3(c Config) (*s3.S3, *session.Session) {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(c.Region),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     c.AccessKeyID,
			SecretAccessKey: c.SecretAccessKey,
		}),
	}))

	s3Session := s3.New(session)

	return s3Session, session
}

type UploadOption struct {
	Bucket string
	Path   string
	File   io.Reader
	ACL    string
}

type UploadResult struct {
	FullPath string
	Filename string
}

func UploadFile(sess *session.Session, u UploadOption) (UploadResult, error) {

	if u.ACL == "" {
		u.ACL = "public-read"
	}

	uploader := s3manager.NewUploader(sess)

	filename := primitive.NewObjectID().Hex() + ".jpg"
	key := u.Path + filename

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(key),
		Body:   u.File,
		ACL:    aws.String(u.ACL),
	})

	var result UploadResult
	result.FullPath = key
	result.Filename = filename

	return result, err
}
