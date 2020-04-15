// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafXssMatchSet(client *Client) ([]Resource, error) {
	req := client.wafconn.ListXssMatchSetsRequest(&waf.ListXssMatchSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.XssMatchSets) > 0 {
		for _, r := range resp.XssMatchSets {

			result = append(result, Resource{
				Type: "aws_waf_xss_match_set",
				ID:   *r.XssMatchSetId,
			})
		}
	}

	return result, nil
}
