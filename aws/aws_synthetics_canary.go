// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/synthetics"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSyntheticsCanary(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := synthetics.NewDescribeCanariesPaginator(client.Syntheticsconn, &synthetics.DescribeCanariesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Canaries {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, terraform.Resource{
				Type:      "aws_synthetics_canary",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	return result, nil
}
