// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

func ListDevicefarmProject(client *Client) {
	req := client.devicefarmconn.ListProjectsRequest(&devicefarm.ListProjectsInput{})

	p := devicefarm.NewListProjectsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_devicefarm_project:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Projects {
			fmt.Println(*r.Arn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_devicefarm_project: %s", err)
	}

}
