// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmPatchBaseline(client *Client) error {
	req := client.ssmconn.DescribePatchBaselinesRequest(&ssm.DescribePatchBaselinesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.BaselineIdentities) > 0 {
		for _, r := range resp.BaselineIdentities {
			fmt.Println(*r.BaselineId)

		}
	}

	return nil
}
