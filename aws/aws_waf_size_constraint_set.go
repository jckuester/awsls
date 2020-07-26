// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafSizeConstraintSet(client *Client) ([]Resource, error) {
	req := client.Wafconn.ListSizeConstraintSetsRequest(&waf.ListSizeConstraintSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SizeConstraintSets) > 0 {
		for _, r := range resp.SizeConstraintSets {

			result = append(result, Resource{
				Type:    "aws_waf_size_constraint_set",
				ID:      *r.SizeConstraintSetId,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	return result, nil
}
