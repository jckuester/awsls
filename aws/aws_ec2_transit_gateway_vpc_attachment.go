// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGatewayVpcAttachment(client *Client) {
	req := client.ec2conn.DescribeTransitGatewayVpcAttachmentsRequest(&ec2.DescribeTransitGatewayVpcAttachmentsInput{})

	p := ec2.NewDescribeTransitGatewayVpcAttachmentsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ec2_transit_gateway_vpc_attachment:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TransitGatewayVpcAttachments {
			fmt.Println(*r.TransitGatewayAttachmentId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ec2_transit_gateway_vpc_attachment: %s", err)
	}

}
