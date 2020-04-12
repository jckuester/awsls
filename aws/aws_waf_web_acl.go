// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafWebAcl(client *Client) {
	req := client.wafconn.ListWebACLsRequest(&waf.ListWebACLsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_waf_web_acl: %s", err)
	} else {
		if len(resp.WebACLs) > 0 {
			fmt.Println("")
			fmt.Println("aws_waf_web_acl:")
			for _, r := range resp.WebACLs {
				fmt.Println(*r.WebACLId)

			}
		}
	}

}
