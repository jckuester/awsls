// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafSqlInjectionMatchSet(client *Client) {
	req := client.wafconn.ListSqlInjectionMatchSetsRequest(&waf.ListSqlInjectionMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_waf_sql_injection_match_set: %s", err)
	} else {
		if len(resp.SqlInjectionMatchSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_waf_sql_injection_match_set:")
			for _, r := range resp.SqlInjectionMatchSets {
				fmt.Println(*r.SqlInjectionMatchSetId)

			}
		}
	}

}
