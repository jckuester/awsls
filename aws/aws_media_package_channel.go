// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
)

func ListMediaPackageChannel(client *Client) ([]Resource, error) {
	req := client.mediapackageconn.ListChannelsRequest(&mediapackage.ListChannelsInput{})

	var result []Resource

	p := mediapackage.NewListChannelsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Channels {
			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type: "aws_media_package_channel",
				ID:   *r.Id,
				Tags: tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
