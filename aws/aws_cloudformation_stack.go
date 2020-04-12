// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func ListCloudformationStack(client *Client) {
	req := client.cloudformationconn.DescribeStacksRequest(&cloudformation.DescribeStacksInput{})

	p := cloudformation.NewDescribeStacksPaginator(req)
	fmt.Println("")
	fmt.Println("aws_cloudformation_stack:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Stacks {
			fmt.Println(*r.StackId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_cloudformation_stack: %s", err)
	}

}
