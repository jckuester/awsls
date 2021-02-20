// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGlueTrigger(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Glueconn.GetTriggersRequest(&glue.GetTriggersInput{})

	var result []terraform.Resource

	p := glue.NewGetTriggersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Triggers {

			result = append(result, terraform.Resource{
				Type:      "aws_glue_trigger",
				ID:        *r.Name,
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
