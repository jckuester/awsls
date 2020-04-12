// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalXssMatchSet(client *Client) error {
	req := client.wafregionalconn.ListXssMatchSetsRequest(&wafregional.ListXssMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.XssMatchSets) > 0 {
		for _, r := range resp.XssMatchSets {
			fmt.Println(*r.XssMatchSetId)
		}
	}

	return nil
}
