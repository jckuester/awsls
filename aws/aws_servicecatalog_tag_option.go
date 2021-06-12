// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListServicecatalogTagOption(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := servicecatalog.NewListTagOptionsPaginator(client.Servicecatalogconn, &servicecatalog.ListTagOptionsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.TagOptionDetails {

			result = append(result, terraform.Resource{
				Type:      "aws_servicecatalog_tag_option",
				ID:        *r.Id,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
