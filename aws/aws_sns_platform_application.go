// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsPlatformApplication(client *Client) ([]Resource, error) {
	req := client.Snsconn.ListPlatformApplicationsRequest(&sns.ListPlatformApplicationsInput{})

	var result []Resource

	p := sns.NewListPlatformApplicationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.PlatformApplications {

			result = append(result, Resource{
				Type:      "aws_sns_platform_application",
				ID:        *r.PlatformApplicationArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
