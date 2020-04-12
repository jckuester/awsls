// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkEnvironment(client *Client) error {
	req := client.elasticbeanstalkconn.DescribeEnvironmentsRequest(&elasticbeanstalk.DescribeEnvironmentsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Environments) > 0 {
		for _, r := range resp.Environments {
			fmt.Println(*r.EnvironmentId)
		}
	}

	return nil
}
