// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbProxy(client *Client) ([]Resource, error) {
	req := client.Rdsconn.DescribeDBProxiesRequest(&rds.DescribeDBProxiesInput{})

	var result []Resource

	p := rds.NewDescribeDBProxiesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBProxies {

			result = append(result, Resource{
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
