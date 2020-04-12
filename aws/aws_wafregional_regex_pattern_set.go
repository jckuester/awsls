// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRegexPatternSet(client *Client) error {
	req := client.wafregionalconn.ListRegexPatternSetsRequest(&wafregional.ListRegexPatternSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.RegexPatternSets) > 0 {
		for _, r := range resp.RegexPatternSets {
			fmt.Println(*r.RegexPatternSetId)

		}
	}

	return nil
}
