package helper

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func SellingPrice(price float64, discount float64) float64 {

	return price - price*discount/100
}

func AddImageToS3(image *multipart.FileHeader) (string, error) {

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		return "", err
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	uploader := manager.NewUploader(s3Client)
	fmt.Println(image.Filename)
	file, err := image.Open()
	if err != nil {

		return "", err
	}
	defer file.Close()
	upload, err1 := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("fashionstore"),
		Key:    aws.String(image.Filename),
		Body:   file,
		ACL:    "public-read",
	})

	if err1 != nil {
		return "", err1
	}

	return upload.Location, nil
}
