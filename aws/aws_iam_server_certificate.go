// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListIamServerCertificate(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Iamconn.ListServerCertificatesRequest(&iam.ListServerCertificatesInput{})

	var result []terraform.Resource

	p := iam.NewListServerCertificatesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ServerCertificateMetadataList {

			result = append(result, terraform.Resource{
				Type:      "aws_iam_server_certificate",
				ID:        *r.ServerCertificateId,
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
