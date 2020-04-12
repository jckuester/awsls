// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2CapacityReservation(client *Client) error {
	req := client.ec2conn.DescribeCapacityReservationsRequest(&ec2.DescribeCapacityReservationsInput{})

	p := ec2.NewDescribeCapacityReservationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.CapacityReservations {
			fmt.Println(*r.CapacityReservationId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.CreateDate)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
