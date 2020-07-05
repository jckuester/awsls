// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailStaticIp(client *Client) ([]Resource, error) {
	req := client.Lightsailconn.GetStaticIpsRequest(&lightsail.GetStaticIpsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.StaticIps) > 0 {
		for _, r := range resp.StaticIps {

			result = append(result, Resource{
				Type:   "aws_lightsail_static_ip",
				ID:     *r.Name,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
