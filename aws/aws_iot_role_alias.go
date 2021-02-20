// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListIotRoleAlias(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Iotconn.ListRoleAliasesRequest(&iot.ListRoleAliasesInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RoleAliases) > 0 {

		for _, r := range resp.RoleAliases {

			result = append(result, terraform.Resource{
				Type:      "aws_iot_role_alias",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
