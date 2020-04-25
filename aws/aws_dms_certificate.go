// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsCertificate(client *Client) ([]Resource, error) {
	req := client.databasemigrationserviceconn.DescribeCertificatesRequest(&databasemigrationservice.DescribeCertificatesInput{})

	var result []Resource

	p := databasemigrationservice.NewDescribeCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Certificates {

			result = append(result, Resource{
				Type:   "aws_dms_certificate",
				ID:     *r.CertificateIdentifier,
				Region: client.databasemigrationserviceconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
