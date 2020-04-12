// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNetworkInterface(client *Client) {
	req := client.ec2conn.DescribeNetworkInterfacesRequest(&ec2.DescribeNetworkInterfacesInput{})

	p := ec2.NewDescribeNetworkInterfacesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_network_interface:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NetworkInterfaces {
			fmt.Println(*r.NetworkInterfaceId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_network_interface: %s", err)
	}

}
