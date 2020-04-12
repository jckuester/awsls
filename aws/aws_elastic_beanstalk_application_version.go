// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkApplicationVersion(client *Client) {
	req := client.elasticbeanstalkconn.DescribeApplicationVersionsRequest(&elasticbeanstalk.DescribeApplicationVersionsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_elastic_beanstalk_application_version: %s", err)
	} else {
		if len(resp.ApplicationVersions) > 0 {
			fmt.Println("")
			fmt.Println("aws_elastic_beanstalk_application_version:")
			for _, r := range resp.ApplicationVersions {
				fmt.Println(*r.ApplicationName)

			}
		}
	}

}
