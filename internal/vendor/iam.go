package vendor

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type AWSIAMUserClient struct {
	iam iam.Client
}

func NewIAMUserClien(cfg aws.Config) AWSIAMUserClient {
	return AWSIAMUserClient{
		iam: *iam.NewFromConfig(cfg),
	}
}

func (aw AWSIAMUserClient) List() ([]string, error) {
	resp, err := aw.iam.ListUsers(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		return nil, err
	}

	userLists := []string{}
	for _, user := range resp.Users {
		userLists = append(userLists, *user.UserName)
	}

	return userLists, nil
}

func (aw AWSIAMUserClient) Search(prefix string) ([]string, error) {
	return nil, nil
}
