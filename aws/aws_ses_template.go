// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSesTemplate(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Sesconn.ListTemplates(ctx, &ses.ListTemplatesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.TemplatesMetadata) > 0 {

		for _, r := range resp.TemplatesMetadata {

			result = append(result, terraform.Resource{
				Type:      "aws_ses_template",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
