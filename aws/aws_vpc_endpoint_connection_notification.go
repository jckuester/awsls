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
		page := p.CurrentPage()

		for _, r := range page.ConnectionNotificationSet {

			result = append(result, Resource{
				Type:   "aws_vpc_endpoint_connection_notification",
				ID:     *r.ConnectionNotificationId,
				Region: client.Ec2conn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
