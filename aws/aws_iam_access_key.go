// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamAccessKey(client *Client) ([]Resource, error) {
	req := client.iamconn.ListAccessKeysRequest(&iam.ListAccessKeysInput{})

	var result []Resource

	p := iam.NewListAccessKeysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.AccessKeyMetadata {

			t := *r.CreateDate
			result = append(result, Resource{
				Type: "aws_iam_access_key",
				ID:   *r.AccessKeyId,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
