// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGuarddutyDetector(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Guarddutyconn.ListDetectorsRequest(&guardduty.ListDetectorsInput{})

	var result []terraform.Resource

	p := guardduty.NewListDetectorsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DetectorIds {

			result = append(result, terraform.Resource{
				Type:      "aws_guardduty_detector",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
