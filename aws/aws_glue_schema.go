// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGlueSchema(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := glue.NewListSchemasPaginator(client.Glueconn, &glue.ListSchemasInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Schemas {

			t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.CreatedTime)
			if err != nil {
				return nil, err
			}
			result = append(result, terraform.Resource{
				Type:      "aws_glue_schema",
				ID:        *r.SchemaArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
