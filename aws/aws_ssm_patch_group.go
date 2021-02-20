// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSsmPatchGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ssmconn.DescribePatchGroupsRequest(&ssm.DescribePatchGroupsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Mappings) > 0 {

		for _, r := range resp.Mappings {

			result = append(result, terraform.Resource{
				Type:      "aws_ssm_patch_group",
				ID:        *r.PatchGroup,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
