// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListAmi(client *Client) {
	req := client.ec2conn.DescribeImagesRequest(&ec2.DescribeImagesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ami: %s", err)
	} else {
		if len(resp.Images) > 0 {
			fmt.Println("")
			fmt.Println("aws_ami:")
			for _, r := range resp.Images {
				fmt.Println(*r.ImageId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
