// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpointConnectionNotification(client *Client) {
	req := client.ec2conn.DescribeVpcEndpointConnectionNotificationsRequest(&ec2.DescribeVpcEndpointConnectionNotificationsInput{})

	p := ec2.NewDescribeVpcEndpointConnectionNotificationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_vpc_endpoint_connection_notification:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ConnectionNotificationSet {
			fmt.Println(*r.ConnectionNotificationId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_vpc_endpoint_connection_notification: %s", err)
	}

}
