// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAmi(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Ec2conn.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Owners: []string{"self"},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Images) > 0 {

		for _, r := range resp.Images {
			if *r.OwnerId != client.AccountID {
				continue
			}
			t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.CreationDate)
			if err != nil {
				return nil, err
			}
			result = append(result, terraform.Resource{
				Type:      "aws_ami",
				ID:        *r.ImageId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
