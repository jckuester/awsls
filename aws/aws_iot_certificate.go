// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotCertificate(client *Client) error {
	req := client.iotconn.ListCertificatesRequest(&iot.ListCertificatesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Certificates) > 0 {
		for _, r := range resp.Certificates {
			fmt.Println(*r.CertificateId)
		}
	}

	return nil
}
