// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2CapacityReservation(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeCapacityReservationsRequest(&ec2.DescribeCapacityReservationsInput{})

	var result []Resource

	p := ec2.NewDescribeCapacityReservationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.CapacityReservations {
			if *r.OwnerId != client.accountid {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreateDate
			result = append(result, Resource{
				Type:      "aws_ec2_capacity_reservation",
				ID:        *r.CapacityReservationId,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
