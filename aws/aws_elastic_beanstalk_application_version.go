// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkApplicationVersion(client *Client) ([]Resource, error) {
	req := client.elasticbeanstalkconn.DescribeApplicationVersionsRequest(&elasticbeanstalk.DescribeApplicationVersionsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ApplicationVersions) > 0 {
		for _, r := range resp.ApplicationVersions {

			result = append(result, Resource{
				Type:   "aws_elastic_beanstalk_application_version",
				ID:     *r.ApplicationName,
				Region: client.elasticbeanstalkconn.Config.Region,
			})
		}
	}

	return result, nil
}
