// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafXssMatchSet(client *Client) {
	req := client.wafconn.ListXssMatchSetsRequest(&waf.ListXssMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_waf_xss_match_set: %s", err)
	} else {
		if len(resp.XssMatchSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_waf_xss_match_set:")
			for _, r := range resp.XssMatchSets {
				fmt.Println(*r.XssMatchSetId)

			}
		}
	}

}
