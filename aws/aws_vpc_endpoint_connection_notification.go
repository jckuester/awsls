// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpointConnectionNotification(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeVpcEndpointConnectionNotificationsRequest(&ec2.DescribeVpcEndpointConnectionNotificationsInput{})

	var result []Resource

	p := ec2.NewDescribeVpcEndpointConnectionNotificationsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ConnectionNotificationSet {

			result = append(result, Resource{
				Type:      "aws_vpc_endpoint_connection_notification",
				ID:        *r.ConnectionNotificationId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
