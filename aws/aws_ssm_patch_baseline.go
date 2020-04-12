// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmPatchBaseline(client *Client) {
	req := client.ssmconn.DescribePatchBaselinesRequest(&ssm.DescribePatchBaselinesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ssm_patch_baseline: %s", err)
	} else {
		if len(resp.BaselineIdentities) > 0 {
			fmt.Println("")
			fmt.Println("aws_ssm_patch_baseline:")
			for _, r := range resp.BaselineIdentities {
				fmt.Println(*r.BaselineId)

			}
		}
	}

}
