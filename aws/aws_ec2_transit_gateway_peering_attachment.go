// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGatewayPeeringAttachment(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeTransitGatewayPeeringAttachmentsRequest(&ec2.DescribeTransitGatewayPeeringAttachmentsInput{})

	var result []Resource

	p := ec2.NewDescribeTransitGatewayPeeringAttachmentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TransitGatewayPeeringAttachments {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_ec2_transit_gateway_peering_attachment",
				ID:        *r.TransitGatewayAttachmentId,
				Region:    client.Region,
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
