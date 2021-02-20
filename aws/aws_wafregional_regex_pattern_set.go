// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafregionalRegexPatternSet(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Wafregionalconn.ListRegexPatternSetsRequest(&wafregional.ListRegexPatternSetsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RegexPatternSets) > 0 {

		for _, r := range resp.RegexPatternSets {

			result = append(result, terraform.Resource{
				Type:      "aws_wafregional_regex_pattern_set",
				ID:        *r.RegexPatternSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
