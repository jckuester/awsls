// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func ListDaxParameterGroup(client *Client) {
	req := client.daxconn.DescribeParameterGroupsRequest(&dax.DescribeParameterGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_dax_parameter_group: %s", err)
	} else {
		if len(resp.ParameterGroups) > 0 {
			fmt.Println("")
			fmt.Println("aws_dax_parameter_group:")
			for _, r := range resp.ParameterGroups {
				fmt.Println(*r.ParameterGroupName)

			}
		}
	}

}
