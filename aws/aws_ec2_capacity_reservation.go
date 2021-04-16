// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEc2CapacityReservation(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeCapacityReservationsPaginator(client.Ec2conn, &ec2.DescribeCapacityReservationsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.CapacityReservations {
			if *r.OwnerId != client.AccountID {
				continue
			}
			t := *r.CreateDate
			result = append(result, terraform.Resource{
				Type:      "aws_ec2_capacity_reservation",
				ID:        *r.CapacityReservationId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
