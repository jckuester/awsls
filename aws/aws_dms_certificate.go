// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsCertificate(client *Client) error {
	req := client.databasemigrationserviceconn.DescribeCertificatesRequest(&databasemigrationservice.DescribeCertificatesInput{})

	p := databasemigrationservice.NewDescribeCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Certificates {
			fmt.Println(*r.CertificateIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
