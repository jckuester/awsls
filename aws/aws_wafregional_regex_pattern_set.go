// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRegexPatternSet(client *Client) ([]Resource, error) {
	req := client.Wafregionalconn.ListRegexPatternSetsRequest(&wafregional.ListRegexPatternSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RegexPatternSets) > 0 {
		for _, r := range resp.RegexPatternSets {

			result = append(result, Resource{
				Type:   "aws_wafregional_regex_pattern_set",
				ID:     *r.RegexPatternSetId,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
