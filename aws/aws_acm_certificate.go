// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func ListAcmCertificate(client *Client) error {
	req := client.acmconn.ListCertificatesRequest(&acm.ListCertificatesInput{})

	p := acm.NewListCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.CertificateSummaryList {
			fmt.Println(*r.CertificateArn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
