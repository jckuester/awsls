// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudformationStackSet(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Cloudformationconn.ListStackSetsRequest(&cloudformation.ListStackSetsInput{})

	var result []terraform.Resource

	p := cloudformation.NewListStackSetsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Summaries {

			result = append(result, terraform.Resource{
				Type:      "aws_cloudformation_stack_set",
				ID:        *r.StackSetName,
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
