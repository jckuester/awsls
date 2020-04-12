// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalSizeConstraintSet(client *Client) error {
	req := client.wafregionalconn.ListSizeConstraintSetsRequest(&wafregional.ListSizeConstraintSetsInput{})

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
