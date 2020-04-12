// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamServerCertificate(client *Client) {
	req := client.iamconn.ListServerCertificatesRequest(&iam.ListServerCertificatesInput{})

	p := iam.NewListServerCertificatesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_iam_server_certificate:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ServerCertificateMetadataList {
			fmt.Println(*r.ServerCertificateId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_iam_server_certificate: %s", err)
	}

}
