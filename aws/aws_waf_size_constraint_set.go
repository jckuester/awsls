// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafSizeConstraintSet(client *Client) {
	req := client.wafconn.ListSizeConstraintSetsRequest(&waf.ListSizeConstraintSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_waf_size_constraint_set: %s", err)
	} else {
		if len(resp.SizeConstraintSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_waf_size_constraint_set:")
			for _, r := range resp.SizeConstraintSets {
				fmt.Println(*r.SizeConstraintSetId)

			}
		}
	}

}
