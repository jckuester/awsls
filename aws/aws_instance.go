// This resource is currently manually added, because the pattern of how
// it's code must be generated differs from other resources (nested for-loop
// of instances inside reservations)

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListInstance(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeInstancesPaginator(client.Ec2conn, &ec2.DescribeInstancesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, reservations := range resp.Reservations {
			if *reservations.OwnerId != client.AccountID {
				continue
			}
			for _, r := range reservations.Instances {

				tags := map[string]string{}
				for _, t := range r.Tags {
					tags[*t.Key] = *t.Value
				}

				t := *r.LaunchTime
				result = append(result, terraform.Resource{
					Type:      "aws_instance",
					ID:        *r.InstanceId,
					Region:    client.Region,
					Profile:   client.Profile,
					AccountID: client.AccountID,
					Tags:      tags,
					CreatedAt: &t,
				})
			}
		}
	}

	return result, nil
}
