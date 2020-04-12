// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSubnetGroup(client *Client) {
	req := client.rdsconn.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{})

	p := rds.NewDescribeDBSubnetGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_db_subnet_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBSubnetGroups {
			fmt.Println(*r.DBSubnetGroupName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_db_subnet_group: %s", err)
	}

}
