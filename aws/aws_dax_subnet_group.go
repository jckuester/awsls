// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDaxSubnetGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Daxconn.DescribeSubnetGroupsRequest(&dax.DescribeSubnetGroupsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SubnetGroups) > 0 {

		for _, r := range resp.SubnetGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_dax_subnet_group",
				ID:        *r.SubnetGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
