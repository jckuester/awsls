// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmPatchBaseline(client *Client) ([]Resource, error) {
	req := client.Ssmconn.DescribePatchBaselinesRequest(&ssm.DescribePatchBaselinesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.BaselineIdentities) > 0 {
		for _, r := range resp.BaselineIdentities {

			result = append(result, Resource{
				Type:   "aws_ssm_patch_baseline",
				ID:     *r.BaselineId,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
