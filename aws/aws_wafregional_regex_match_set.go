// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRegexMatchSet(client *Client) {
	req := client.wafregionalconn.ListRegexMatchSetsRequest(&wafregional.ListRegexMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_regex_match_set: %s", err)
	} else {
		if len(resp.RegexMatchSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_regex_match_set:")
			for _, r := range resp.RegexMatchSets {
				fmt.Println(*r.RegexMatchSetId)

			}
		}
	}

}
