// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpointConnectionNotification(client *Client) error {
	req := client.ec2conn.DescribeVpcEndpointConnectionNotificationsRequest(&ec2.DescribeVpcEndpointConnectionNotificationsInput{})

	p := ec2.NewDescribeVpcEndpointConnectionNotificationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ConnectionNotificationSet {
			fmt.Println(*r.ConnectionNotificationId)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
