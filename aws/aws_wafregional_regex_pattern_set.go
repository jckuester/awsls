// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRegexPatternSet(client *Client) {
	req := client.wafregionalconn.ListRegexPatternSetsRequest(&wafregional.ListRegexPatternSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_regex_pattern_set: %s", err)
	} else {
		if len(resp.RegexPatternSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_regex_pattern_set:")
			for _, r := range resp.RegexPatternSets {
				fmt.Println(*r.RegexPatternSetId)

			}
		}
	}

}
