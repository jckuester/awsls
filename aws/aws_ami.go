// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListAmi(client *Client) error {
	req := client.ec2conn.DescribeImagesRequest(&ec2.DescribeImagesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Images) > 0 {
		for _, r := range resp.Images {
			fmt.Println(*r.ImageId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
