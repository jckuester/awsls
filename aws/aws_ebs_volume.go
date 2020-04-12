// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEbsVolume(client *Client) error {
	req := client.ec2conn.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})

	p := ec2.NewDescribeVolumesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Volumes {
			fmt.Println(*r.VolumeId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
