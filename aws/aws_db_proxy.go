// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDbProxy(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Rdsconn.DescribeDBProxiesRequest(&rds.DescribeDBProxiesInput{})

	var result []terraform.Resource

	p := rds.NewDescribeDBProxiesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DBProxies {

			result = append(result, terraform.Resource{
				Type:      "aws_db_proxy",
				ID:        *r.DBProxyName,
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
