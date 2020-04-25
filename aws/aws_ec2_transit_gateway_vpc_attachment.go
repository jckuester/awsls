// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGatewayVpcAttachment(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeTransitGatewayVpcAttachmentsRequest(&ec2.DescribeTransitGatewayVpcAttachmentsInput{})

	var result []Resource

	p := ec2.NewDescribeTransitGatewayVpcAttachmentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TransitGatewayVpcAttachments {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_ec2_transit_gateway_vpc_attachment",
				ID:        *r.TransitGatewayAttachmentId,
				Region:    client.ec2conn.Config.Region,
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
