// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRegexMatchSet(client *Client) ([]Resource, error) {
	req := client.wafregionalconn.ListRegexMatchSetsRequest(&wafregional.ListRegexMatchSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RegexMatchSets) > 0 {
		for _, r := range resp.RegexMatchSets {

			result = append(result, Resource{
				Type: "aws_wafregional_regex_match_set",
				ID:   *r.RegexMatchSetId,
			})
		}
	}

	return result, nil
}
