// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListMediaStoreContainer(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Mediastoreconn.ListContainersRequest(&mediastore.ListContainersInput{})

	var result []terraform.Resource

	p := mediastore.NewListContainersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Containers {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_media_store_container",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
