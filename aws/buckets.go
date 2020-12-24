package buckets

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func WriteToBucket(filename string) {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open file %q. %v", filename, err)
	}

	result, awsErr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(""),
		Key:    aws.String(""),
		Body:   f,
	})

	if awsErr != nil {
		fmt.Errorf("failed to upload file %v", err)
	}
	//fmt.Printf("file uploaded to %s\n". result.Location)
	fmt.Println(result)
}
