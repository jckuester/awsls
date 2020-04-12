// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNetworkAcl(client *Client) {
	req := client.ec2conn.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})

	p := ec2.NewDescribeNetworkAclsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_network_acl:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NetworkAcls {
			fmt.Println(*r.NetworkAclId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_network_acl: %s", err)
	}

}
