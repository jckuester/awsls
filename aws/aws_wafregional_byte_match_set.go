// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalByteMatchSet(client *Client) ([]Resource, error) {
	req := client.Wafregionalconn.ListByteMatchSetsRequest(&wafregional.ListByteMatchSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ByteMatchSets) > 0 {
		for _, r := range resp.ByteMatchSets {

			result = append(result, Resource{
				Type:   "aws_wafregional_byte_match_set",
				ID:     *r.ByteMatchSetId,
				Region: client.Wafregionalconn.Config.Region,
			})
		}
	}

	return result, nil
}
