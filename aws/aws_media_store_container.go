// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

func ListMediaStoreContainer(client *Client) error {
	req := client.mediastoreconn.ListContainersRequest(&mediastore.ListContainersInput{})

	p := mediastore.NewListContainersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Containers {
			fmt.Println(*r.Name)

			fmt.Printf("CreatedAt: %s\n", *r.CreationTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
