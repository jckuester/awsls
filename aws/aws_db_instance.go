// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbInstance(client *Client) {
	req := client.rdsconn.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})

	p := rds.NewDescribeDBInstancesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_db_instance:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBInstances {
			fmt.Println(*r.DBInstanceIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_db_instance: %s", err)
	}

}
