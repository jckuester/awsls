// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalSizeConstraintSet(client *Client) {
	req := client.wafregionalconn.ListSizeConstraintSetsRequest(&wafregional.ListSizeConstraintSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_size_constraint_set: %s", err)
	} else {
		if len(resp.SizeConstraintSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_size_constraint_set:")
			for _, r := range resp.SizeConstraintSets {
				fmt.Println(*r.SizeConstraintSetId)

			}
		}
	}

}
