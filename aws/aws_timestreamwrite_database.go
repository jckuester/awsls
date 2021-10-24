// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListTimestreamwriteDatabase(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := timestreamwrite.NewListDatabasesPaginator(client.Timestreamwriteconn, &timestreamwrite.ListDatabasesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Databases {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_timestreamwrite_database",
				ID:        *r.DatabaseName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
