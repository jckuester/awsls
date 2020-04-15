// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotCertificate(client *Client) ([]Resource, error) {
	req := client.iotconn.ListCertificatesRequest(&iot.ListCertificatesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Certificates) > 0 {
		for _, r := range resp.Certificates {

			t := *r.CreationDate
			result = append(result, Resource{
				Type: "aws_iot_certificate",
				ID:   *r.CertificateId,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
