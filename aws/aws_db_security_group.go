// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSecurityGroup(client *Client) {
	req := client.rdsconn.DescribeDBSecurityGroupsRequest(&rds.DescribeDBSecurityGroupsInput{})

	p := rds.NewDescribeDBSecurityGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_db_security_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBSecurityGroups {
			fmt.Println(*r.DBSecurityGroupName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_db_security_group: %s", err)
	}

}
