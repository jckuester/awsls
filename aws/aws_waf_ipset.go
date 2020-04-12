// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafIpset(client *Client) error {
	req := client.wafconn.ListIPSetsRequest(&waf.ListIPSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.IPSets) > 0 {
		for _, r := range resp.IPSets {
			fmt.Println(*r.IPSetId)

		}
	}

	return nil
}
