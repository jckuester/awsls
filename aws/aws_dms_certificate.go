// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDmsCertificate(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := databasemigrationservice.NewDescribeCertificatesPaginator(client.Databasemigrationserviceconn, &databasemigrationservice.DescribeCertificatesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Certificates {

			result = append(result, terraform.Resource{
				Type:      "aws_dms_certificate",
				ID:        *r.CertificateIdentifier,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
