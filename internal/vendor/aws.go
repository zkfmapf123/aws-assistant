package vendor

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func NewAWSConfig() (aws.Config, error) {
	acKey := os.Getenv("AWS_ACCESS_KEY_ID")
	seKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	if acKey == "" || seKey == "" || region == "" {
		return aws.Config{}, fmt.Errorf("AWS credentials not set")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(acKey, seKey, ""),
		))
	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}

// func newAWSConfigisTest() (aws.Config, error) {
// 	cfg, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		return aws.Config{}, nil
// 	}

// 	return cfg, nil
// }
