// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmMaintenanceWindow(client *Client) ([]Resource, error) {
	req := client.Ssmconn.DescribeMaintenanceWindowsRequest(&ssm.DescribeMaintenanceWindowsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.WindowIdentities) > 0 {
		for _, r := range resp.WindowIdentities {

			result = append(result, Resource{
				Type:    "aws_ssm_maintenance_window",
				ID:      *r.WindowId,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	return result, nil
}
