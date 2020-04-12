// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/athena"
)

func ListAthenaWorkgroup(client *Client) {
	req := client.athenaconn.ListWorkGroupsRequest(&athena.ListWorkGroupsInput{})

	p := athena.NewListWorkGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_athena_workgroup:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.WorkGroups {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_athena_workgroup: %s", err)
	}

}
