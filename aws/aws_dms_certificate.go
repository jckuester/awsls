// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsCertificate(client *Client) {
	req := client.databasemigrationserviceconn.DescribeCertificatesRequest(&databasemigrationservice.DescribeCertificatesInput{})

	p := databasemigrationservice.NewDescribeCertificatesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_dms_certificate:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Certificates {
			fmt.Println(*r.CertificateIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_dms_certificate: %s", err)
	}

}
