// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListFmsPolicy(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Fmsconn.ListPoliciesRequest(&fms.ListPoliciesInput{})

	var result []terraform.Resource

	p := fms.NewListPoliciesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.PolicyList {

			result = append(result, terraform.Resource{
				Type:      "aws_fms_policy",
				ID:        *r.PolicyId,
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
