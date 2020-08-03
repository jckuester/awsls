// TODO this resource is currently manually added, because the pattern of how
// it's code must be generated differs from all other resources (nested for-loop
// of instances inside reservations)

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListInstance(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})

	var result []Resource

	p := ec2.NewDescribeInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, reservations := range page.Reservations {
			if *reservations.OwnerId != client.AccountID {
				continue
			}
			for _, r := range reservations.Instances {

				tags := map[string]string{}
				for _, t := range r.Tags {
					tags[*t.Key] = *t.Value
				}

				t := *r.LaunchTime
				result = append(result, Resource{
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

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
