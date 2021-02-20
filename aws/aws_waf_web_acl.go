// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafWebAcl(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Wafconn.ListWebACLsRequest(&waf.ListWebACLsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.WebACLs) > 0 {

		for _, r := range resp.WebACLs {

			result = append(result, terraform.Resource{
				Type:      "aws_waf_web_acl",
				ID:        *r.WebACLId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
