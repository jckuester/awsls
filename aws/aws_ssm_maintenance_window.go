// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmMaintenanceWindow(client *Client) {
	req := client.ssmconn.DescribeMaintenanceWindowsRequest(&ssm.DescribeMaintenanceWindowsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ssm_maintenance_window: %s", err)
	} else {
		if len(resp.WindowIdentities) > 0 {
			fmt.Println("")
			fmt.Println("aws_ssm_maintenance_window:")
			for _, r := range resp.WindowIdentities {
				fmt.Println(*r.WindowId)

			}
		}
	}

}
