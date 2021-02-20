// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSsmAssociation(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ssmconn.ListAssociationsRequest(&ssm.ListAssociationsInput{})

	var result []terraform.Resource

	p := ssm.NewListAssociationsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Associations {

			result = append(result, terraform.Resource{
				Type:      "aws_ssm_association",
				ID:        *r.AssociationId,
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
