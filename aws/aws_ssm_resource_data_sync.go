// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSsmResourceDataSync(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ssmconn.ListResourceDataSyncRequest(&ssm.ListResourceDataSyncInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ResourceDataSyncItems) > 0 {

		for _, r := range resp.ResourceDataSyncItems {

			result = append(result, terraform.Resource{
				Type:      "aws_ssm_resource_data_sync",
				ID:        *r.SyncName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
