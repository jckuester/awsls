// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafByteMatchSet(client *Client) error {
	req := client.wafconn.ListByteMatchSetsRequest(&waf.ListByteMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.ByteMatchSets) > 0 {
		for _, r := range resp.ByteMatchSets {
			fmt.Println(*r.ByteMatchSetId)
		}
	}

	return nil
}
