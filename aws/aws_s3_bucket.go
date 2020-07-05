// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListS3Bucket(client *Client) ([]Resource, error) {
	req := client.S3conn.ListBucketsRequest(&s3.ListBucketsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Buckets) > 0 {
		for _, r := range resp.Buckets {

			t := *r.CreationDate
			result = append(result, Resource{
				Type:   "aws_s3_bucket",
				ID:     *r.Name,
				Region: client.S3conn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
