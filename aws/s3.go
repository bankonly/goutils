package aws

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func UploadFile(sess *session.Session, file io.Reader) {
	uploader := s3manager.NewUploader(sess)
	uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("smart-school-storage"),
		Key:    aws.String("images/test/image.png"),
		Body:   file,
		ACL:    aws.String("public-read"),
	})
}
