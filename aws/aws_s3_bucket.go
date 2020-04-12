// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListS3Bucket(client *Client) {
	req := client.s3conn.ListBucketsRequest(&s3.ListBucketsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_s3_bucket: %s", err)
	} else {
		if len(resp.Buckets) > 0 {
			fmt.Println("")
			fmt.Println("aws_s3_bucket:")
			for _, r := range resp.Buckets {
				fmt.Println(*r.Name)

			}
		}
	}

}
