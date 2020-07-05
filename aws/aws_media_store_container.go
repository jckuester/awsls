// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

func ListMediaStoreContainer(client *Client) ([]Resource, error) {
	req := client.Mediastoreconn.ListContainersRequest(&mediastore.ListContainersInput{})

	var result []Resource

	p := mediastore.NewListContainersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Containers {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:   "aws_media_store_container",
				ID:     *r.Name,
				Region: client.Mediastoreconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
