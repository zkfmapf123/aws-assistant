package vendor

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSS3Client struct {
	s3 s3.Client
}

func NewS3Client(cfg aws.Config) AWSS3Client {
	return AWSS3Client{
		s3: *s3.NewFromConfig(cfg),
	}
}

func (aw AWSS3Client) List() ([]string, error) {
	resp, err := aw.s3.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	bucketList := []string{}
	for _, bucket := range resp.Buckets {
		bucketList = append(bucketList, *bucket.Name)
	}

	return bucketList, nil
}

func (aw AWSS3Client) Search(prefix string) ([]string, error) {
	bucketList, err := aw.List()
	if err != nil {
		return nil, err
	}

	searchBucketList := []string{}
	for _, v := range bucketList {
		if strings.Contains(v, prefix) {
			searchBucketList = append(searchBucketList, v)
		}
	}

	return searchBucketList, nil
}
