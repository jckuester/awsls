// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

func ListDevicefarmProject(client *Client) error {
	req := client.devicefarmconn.ListProjectsRequest(&devicefarm.ListProjectsInput{})

	p := devicefarm.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Projects {
			fmt.Println(*r.Arn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
