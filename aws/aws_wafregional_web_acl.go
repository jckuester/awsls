// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalWebAcl(client *Client) {
	req := client.wafregionalconn.ListWebACLsRequest(&wafregional.ListWebACLsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_web_acl: %s", err)
	} else {
		if len(resp.WebACLs) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_web_acl:")
			for _, r := range resp.WebACLs {
				fmt.Println(*r.WebACLId)

			}
		}
	}

}
