// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkApplication(client *Client) error {
	req := client.elasticbeanstalkconn.DescribeApplicationsRequest(&elasticbeanstalk.DescribeApplicationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Applications) > 0 {
		for _, r := range resp.Applications {
			fmt.Println(*r.ApplicationName)

		}
	}

	return nil
}
