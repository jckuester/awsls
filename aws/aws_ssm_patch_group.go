// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmPatchGroup(client *Client) {
	req := client.ssmconn.DescribePatchGroupsRequest(&ssm.DescribePatchGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ssm_patch_group: %s", err)
	} else {
		if len(resp.Mappings) > 0 {
			fmt.Println("")
			fmt.Println("aws_ssm_patch_group:")
			for _, r := range resp.Mappings {
				fmt.Println(*r.PatchGroup)

			}
		}
	}

}
