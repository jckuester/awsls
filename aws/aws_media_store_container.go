// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

func ListMediaStoreContainer(client *Client) {
	req := client.mediastoreconn.ListContainersRequest(&mediastore.ListContainersInput{})

	p := mediastore.NewListContainersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_media_store_container:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Containers {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_media_store_container: %s", err)
	}

}
