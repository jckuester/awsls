// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWorkspacesIpGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Workspacesconn.DescribeIpGroupsRequest(&workspaces.DescribeIpGroupsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Result) > 0 {

		for _, r := range resp.Result {

			result = append(result, terraform.Resource{
				Type:      "aws_workspaces_ip_group",
				ID:        *r.GroupId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
