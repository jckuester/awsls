// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalWebAcl(client *Client) error {
	req := client.wafregionalconn.ListWebACLsRequest(&wafregional.ListWebACLsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.WebACLs) > 0 {
		for _, r := range resp.WebACLs {
			fmt.Println(*r.WebACLId)

		}
	}

	return nil
}
