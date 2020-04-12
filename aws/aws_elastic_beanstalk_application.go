// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkApplication(client *Client) {
	req := client.elasticbeanstalkconn.DescribeApplicationsRequest(&elasticbeanstalk.DescribeApplicationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_elastic_beanstalk_application: %s", err)
	} else {
		if len(resp.Applications) > 0 {
			fmt.Println("")
			fmt.Println("aws_elastic_beanstalk_application:")
			for _, r := range resp.Applications {
				fmt.Println(*r.ApplicationName)

			}
		}
	}

}
