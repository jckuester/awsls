// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func ListAcmCertificate(client *Client) {
	req := client.acmconn.ListCertificatesRequest(&acm.ListCertificatesInput{})

	p := acm.NewListCertificatesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_acm_certificate:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.CertificateSummaryList {
			fmt.Println(*r.CertificateArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_acm_certificate: %s", err)
	}

}
