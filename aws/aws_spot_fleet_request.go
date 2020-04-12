// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListSpotFleetRequest(client *Client) {
	req := client.ec2conn.DescribeSpotFleetRequestsRequest(&ec2.DescribeSpotFleetRequestsInput{})

	p := ec2.NewDescribeSpotFleetRequestsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_spot_fleet_request:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SpotFleetRequestConfigs {
			fmt.Println(*r.SpotFleetRequestId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_spot_fleet_request: %s", err)
	}

}
