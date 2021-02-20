// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDaxParameterGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Daxconn.DescribeParameterGroupsRequest(&dax.DescribeParameterGroupsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ParameterGroups) > 0 {

		for _, r := range resp.ParameterGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_dax_parameter_group",
				ID:        *r.ParameterGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
