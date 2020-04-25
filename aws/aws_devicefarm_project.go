// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

func ListDevicefarmProject(client *Client) ([]Resource, error) {
	req := client.devicefarmconn.ListProjectsRequest(&devicefarm.ListProjectsInput{})

	var result []Resource

	p := devicefarm.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Projects {

			result = append(result, Resource{
				Type:   "aws_devicefarm_project",
				ID:     *r.Arn,
				Region: client.devicefarmconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
