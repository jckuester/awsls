// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWorkspacesIpGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Workspacesconn.DescribeIpGroups(ctx, &workspaces.DescribeIpGroupsInput{})
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
