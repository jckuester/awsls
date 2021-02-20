// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListXrayGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Xrayconn.GetGroupsRequest(&xray.GetGroupsInput{})

	var result []terraform.Resource

	p := xray.NewGetGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Groups {

			result = append(result, terraform.Resource{
				Type:      "aws_xray_group",
				ID:        *r.GroupARN,
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
