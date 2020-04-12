// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafSizeConstraintSet(client *Client) error {
	req := client.wafconn.ListSizeConstraintSetsRequest(&waf.ListSizeConstraintSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.SizeConstraintSets) > 0 {
		for _, r := range resp.SizeConstraintSets {
			fmt.Println(*r.SizeConstraintSetId)

		}
	}

	return nil
}
