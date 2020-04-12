// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func ListDaxSubnetGroup(client *Client) {
	req := client.daxconn.DescribeSubnetGroupsRequest(&dax.DescribeSubnetGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_dax_subnet_group: %s", err)
	} else {
		if len(resp.SubnetGroups) > 0 {
			fmt.Println("")
			fmt.Println("aws_dax_subnet_group:")
			for _, r := range resp.SubnetGroups {
				fmt.Println(*r.SubnetGroupName)

			}
		}
	}

}
