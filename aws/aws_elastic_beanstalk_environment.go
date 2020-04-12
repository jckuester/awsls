// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkEnvironment(client *Client) {
	req := client.elasticbeanstalkconn.DescribeEnvironmentsRequest(&elasticbeanstalk.DescribeEnvironmentsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_elastic_beanstalk_environment: %s", err)
	} else {
		if len(resp.Environments) > 0 {
			fmt.Println("")
			fmt.Println("aws_elastic_beanstalk_environment:")
			for _, r := range resp.Environments {
				fmt.Println(*r.EnvironmentId)

			}
		}
	}

}
