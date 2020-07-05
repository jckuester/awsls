// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafIpset(client *Client) ([]Resource, error) {
	req := client.Wafconn.ListIPSetsRequest(&waf.ListIPSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.IPSets) > 0 {
		for _, r := range resp.IPSets {

			result = append(result, Resource{
				Type:   "aws_waf_ipset",
				ID:     *r.IPSetId,
				Region: client.Wafconn.Config.Region,
			})
		}
	}

	return result, nil
}
