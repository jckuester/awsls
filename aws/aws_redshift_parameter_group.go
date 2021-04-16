// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRedshiftParameterGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := redshift.NewDescribeClusterParameterGroupsPaginator(client.Redshiftconn, &redshift.DescribeClusterParameterGroupsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.ParameterGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_redshift_parameter_group",
				ID:        *r.ParameterGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
