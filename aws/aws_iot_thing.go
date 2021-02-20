// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListIotThing(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Iotconn.ListThingsRequest(&iot.ListThingsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Things) > 0 {

		for _, r := range resp.Things {

			result = append(result, terraform.Resource{
				Type:      "aws_iot_thing",
				ID:        *r.ThingName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
