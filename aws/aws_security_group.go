// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListSecurityGroup(client *Client) {
	req := client.ec2conn.DescribeSecurityGroupsRequest(&ec2.DescribeSecurityGroupsInput{})

	p := ec2.NewDescribeSecurityGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_security_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecurityGroups {
			fmt.Println(*r.GroupId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_security_group: %s", err)
	}

}
