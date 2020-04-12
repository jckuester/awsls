// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkApplicationVersion(client *Client) error {
	req := client.elasticbeanstalkconn.DescribeApplicationVersionsRequest(&elasticbeanstalk.DescribeApplicationVersionsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.ApplicationVersions) > 0 {
		for _, r := range resp.ApplicationVersions {
			fmt.Println(*r.ApplicationName)

		}
	}

	return nil
}
