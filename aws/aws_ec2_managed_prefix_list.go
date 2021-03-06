// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEc2ManagedPrefixList(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ec2conn.DescribeManagedPrefixListsRequest(&ec2.DescribeManagedPrefixListsInput{})

	var result []terraform.Resource

	p := ec2.NewDescribeManagedPrefixListsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.PrefixLists {
			if *r.OwnerId != client.AccountID {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, terraform.Resource{
				Type:      "aws_ec2_managed_prefix_list",
				ID:        *r.PrefixListId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
