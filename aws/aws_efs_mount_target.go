// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func ListEfsMountTarget(client *Client) ([]Resource, error) {
	req := client.efsconn.DescribeMountTargetsRequest(&efs.DescribeMountTargetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.MountTargets) > 0 {
		for _, r := range resp.MountTargets {
			if *r.OwnerId != client.accountid {
				continue
			}

			result = append(result, Resource{
				Type:   "aws_efs_mount_target",
				ID:     *r.MountTargetId,
				Region: client.efsconn.Config.Region,
			})
		}
	}

	return result, nil
}
