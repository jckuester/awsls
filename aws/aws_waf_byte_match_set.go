// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafByteMatchSet(client *Client) ([]Resource, error) {
	req := client.Wafconn.ListByteMatchSetsRequest(&waf.ListByteMatchSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ByteMatchSets) > 0 {
		for _, r := range resp.ByteMatchSets {

			result = append(result, Resource{
				Type:    "aws_waf_byte_match_set",
				ID:      *r.ByteMatchSetId,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	return result, nil
}
