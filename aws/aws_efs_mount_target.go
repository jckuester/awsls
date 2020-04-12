// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func ListEfsMountTarget(client *Client) {
	req := client.efsconn.DescribeMountTargetsRequest(&efs.DescribeMountTargetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_efs_mount_target: %s", err)
	} else {
		if len(resp.MountTargets) > 0 {
			fmt.Println("")
			fmt.Println("aws_efs_mount_target:")
			for _, r := range resp.MountTargets {
				fmt.Println(*r.MountTargetId)

			}
		}
	}

}
