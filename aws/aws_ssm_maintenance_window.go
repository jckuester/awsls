// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmMaintenanceWindow(client *Client) error {
	req := client.ssmconn.DescribeMaintenanceWindowsRequest(&ssm.DescribeMaintenanceWindowsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.WindowIdentities) > 0 {
		for _, r := range resp.WindowIdentities {
			fmt.Println(*r.WindowId)
		}
	}

	return nil
}
