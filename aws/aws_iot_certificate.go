// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListIotCertificate(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Iotconn.ListCertificatesRequest(&iot.ListCertificatesInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Certificates) > 0 {

		for _, r := range resp.Certificates {

			t := *r.CreationDate
			result = append(result, terraform.Resource{
				Type:      "aws_iot_certificate",
				ID:        *r.CertificateId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
