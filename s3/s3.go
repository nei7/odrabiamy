package s3

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nei7/odrabiamy/config"
	"github.com/nei7/odrabiamy/logger"
)

func InitSession() *session.Session {
	session, err := session.NewSession(&aws.Config{Region: aws.String(config.Config.S3.Region)})
	if err != nil {
		log.Fatal(err)
	}

	return session
}

func UploadFile(session *session.Session, dir string) error {
	fi, err := os.Open(dir)
	if err != nil {
		logger.ErrorLogger.Fatalf("S3: %v \n", err)
		return err
	}
	defer fi.Close()

	info, err := fi.Stat()
	if err != nil {
		logger.ErrorLogger.Fatalf("S3: %v \n", err)
		return err
	}

	buf := make([]byte, info.Size())
	fi.Read(buf)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               &config.Config.S3.Bucket,
		Key:                  aws.String(dir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buf),
		ContentLength:        aws.Int64(info.Size()),
		ContentType:          aws.String(http.DetectContentType(buf)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
