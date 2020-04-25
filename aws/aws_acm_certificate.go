// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func ListAcmCertificate(client *Client) ([]Resource, error) {
	req := client.acmconn.ListCertificatesRequest(&acm.ListCertificatesInput{})

	var result []Resource

	p := acm.NewListCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.CertificateSummaryList {

			result = append(result, Resource{
				Type:   "aws_acm_certificate",
				ID:     *r.CertificateArn,
				Region: client.acmconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
