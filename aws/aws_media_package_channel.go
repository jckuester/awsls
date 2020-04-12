// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
)

func ListMediaPackageChannel(client *Client) error {
	req := client.mediapackageconn.ListChannelsRequest(&mediapackage.ListChannelsInput{})

	p := mediapackage.NewListChannelsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Channels {
			fmt.Println(*r.Id)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
