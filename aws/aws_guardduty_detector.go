// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

func ListGuarddutyDetector(client *Client) ([]Resource, error) {
	req := client.Guarddutyconn.ListDetectorsRequest(&guardduty.ListDetectorsInput{})

	var result []Resource

	p := guardduty.NewListDetectorsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DetectorIds {

			result = append(result, Resource{
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
