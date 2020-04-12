// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalByteMatchSet(client *Client) {
	req := client.wafregionalconn.ListByteMatchSetsRequest(&wafregional.ListByteMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_byte_match_set: %s", err)
	} else {
		if len(resp.ByteMatchSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_byte_match_set:")
			for _, r := range resp.ByteMatchSets {
				fmt.Println(*r.ByteMatchSetId)

			}
		}
	}

}
