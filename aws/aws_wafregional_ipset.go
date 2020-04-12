// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalIpset(client *Client) {
	req := client.wafregionalconn.ListIPSetsRequest(&wafregional.ListIPSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_ipset: %s", err)
	} else {
		if len(resp.IPSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_ipset:")
			for _, r := range resp.IPSets {
				fmt.Println(*r.IPSetId)

			}
		}
	}

}