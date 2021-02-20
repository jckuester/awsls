// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAcmCertificate(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Acmconn.ListCertificatesRequest(&acm.ListCertificatesInput{})

	var result []terraform.Resource

	p := acm.NewListCertificatesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.CertificateSummaryList {

			result = append(result, terraform.Resource{
				Type:      "aws_acm_certificate",
				ID:        *r.CertificateArn,
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
