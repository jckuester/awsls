// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func ListWorkspacesIpGroup(client *Client) {
	req := client.workspacesconn.DescribeIpGroupsRequest(&workspaces.DescribeIpGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_workspaces_ip_group: %s", err)
	} else {
		if len(resp.Result) > 0 {
			fmt.Println("")
			fmt.Println("aws_workspaces_ip_group:")
			for _, r := range resp.Result {
				fmt.Println(*r.GroupId)

			}
		}
	}

}
