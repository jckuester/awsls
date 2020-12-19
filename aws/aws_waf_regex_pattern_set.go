// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafRegexPatternSet(client *Client) ([]Resource, error) {
	req := client.Wafconn.ListRegexPatternSetsRequest(&waf.ListRegexPatternSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RegexPatternSets) > 0 {

		for _, r := range resp.RegexPatternSets {

			result = append(result, Resource{
				Type:      "aws_waf_regex_pattern_set",
				ID:        *r.RegexPatternSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
