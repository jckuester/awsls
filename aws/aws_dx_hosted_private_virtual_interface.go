// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxHostedPrivateVirtualInterface(client *Client) {
	req := client.directconnectconn.DescribeVirtualInterfacesRequest(&directconnect.DescribeVirtualInterfacesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_dx_hosted_private_virtual_interface: %s", err)
	} else {
		if len(resp.VirtualInterfaces) > 0 {
			fmt.Println("")
			fmt.Println("aws_dx_hosted_private_virtual_interface:")
			for _, r := range resp.VirtualInterfaces {
				fmt.Println(*r.VirtualInterfaceId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
