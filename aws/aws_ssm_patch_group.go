// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmPatchGroup(client *Client) error {
	req := client.ssmconn.DescribePatchGroupsRequest(&ssm.DescribePatchGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Mappings) > 0 {
		for _, r := range resp.Mappings {
			fmt.Println(*r.PatchGroup)
		}
	}

	return nil
}
