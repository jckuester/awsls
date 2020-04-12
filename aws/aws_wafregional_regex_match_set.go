// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRegexMatchSet(client *Client) error {
	req := client.wafregionalconn.ListRegexMatchSetsRequest(&wafregional.ListRegexMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.RegexMatchSets) > 0 {
		for _, r := range resp.RegexMatchSets {
			fmt.Println(*r.RegexMatchSetId)
		}
	}

	return nil
}
