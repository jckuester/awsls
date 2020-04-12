// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbParameterGroup(client *Client) {
	req := client.rdsconn.DescribeDBParameterGroupsRequest(&rds.DescribeDBParameterGroupsInput{})

	p := rds.NewDescribeDBParameterGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_db_parameter_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBParameterGroups {
			fmt.Println(*r.DBParameterGroupName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_db_parameter_group: %s", err)
	}

}
