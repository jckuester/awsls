// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafByteMatchSet(client *Client) {
	req := client.wafconn.ListByteMatchSetsRequest(&waf.ListByteMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_waf_byte_match_set: %s", err)
	} else {
		if len(resp.ByteMatchSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_waf_byte_match_set:")
			for _, r := range resp.ByteMatchSets {
				fmt.Println(*r.ByteMatchSetId)

			}
		}
	}

}
