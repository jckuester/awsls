// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotCertificate(client *Client) {
	req := client.iotconn.ListCertificatesRequest(&iot.ListCertificatesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_iot_certificate: %s", err)
	} else {
		if len(resp.Certificates) > 0 {
			fmt.Println("")
			fmt.Println("aws_iot_certificate:")
			for _, r := range resp.Certificates {
				fmt.Println(*r.CertificateId)

			}
		}
	}

}
