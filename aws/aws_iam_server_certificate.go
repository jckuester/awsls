// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamServerCertificate(client *Client) error {
	req := client.iamconn.ListServerCertificatesRequest(&iam.ListServerCertificatesInput{})

	p := iam.NewListServerCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ServerCertificateMetadataList {
			fmt.Println(*r.ServerCertificateId)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
