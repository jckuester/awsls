// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamServerCertificate(client *Client) ([]Resource, error) {
	req := client.Iamconn.ListServerCertificatesRequest(&iam.ListServerCertificatesInput{})

	var result []Resource

	p := iam.NewListServerCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ServerCertificateMetadataList {

			result = append(result, Resource{
				Type:    "aws_iam_server_certificate",
				ID:      *r.ServerCertificateId,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
