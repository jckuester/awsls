// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafSqlInjectionMatchSet(client *Client) error {
	req := client.wafconn.ListSqlInjectionMatchSetsRequest(&waf.ListSqlInjectionMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.SqlInjectionMatchSets) > 0 {
		for _, r := range resp.SqlInjectionMatchSets {
			fmt.Println(*r.SqlInjectionMatchSetId)
		}
	}

	return nil
}
