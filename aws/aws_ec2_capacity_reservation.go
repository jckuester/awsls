// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2CapacityReservation(client *Client) {
	req := client.ec2conn.DescribeCapacityReservationsRequest(&ec2.DescribeCapacityReservationsInput{})

	p := ec2.NewDescribeCapacityReservationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ec2_capacity_reservation:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.CapacityReservations {
			fmt.Println(*r.CapacityReservationId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ec2_capacity_reservation: %s", err)
	}

}
