// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListFsxLustreFileSystem(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Fsxconn.DescribeFileSystemsRequest(&fsx.DescribeFileSystemsInput{})

	var result []terraform.Resource

	p := fsx.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.FileSystems {
			if *r.OwnerId != client.AccountID {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_fsx_lustre_file_system",
				ID:        *r.FileSystemId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
