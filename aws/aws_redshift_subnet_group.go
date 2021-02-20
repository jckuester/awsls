// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRedshiftSubnetGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Redshiftconn.DescribeClusterSubnetGroupsRequest(&redshift.DescribeClusterSubnetGroupsInput{})

	var result []terraform.Resource

	p := redshift.NewDescribeClusterSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ClusterSubnetGroups {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, terraform.Resource{
				Type:      "aws_redshift_subnet_group",
				ID:        *r.ClusterSubnetGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
