// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmResourceDataSync(client *Client) ([]Resource, error) {
	req := client.Ssmconn.ListResourceDataSyncRequest(&ssm.ListResourceDataSyncInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ResourceDataSyncItems) > 0 {
		for _, r := range resp.ResourceDataSyncItems {

			result = append(result, Resource{
				Type:   "aws_ssm_resource_data_sync",
				ID:     *r.SyncName,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
