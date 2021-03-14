// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListFsxWindowsFileSystem(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := fsx.NewDescribeFileSystemsPaginator(client.Fsxconn, &fsx.DescribeFileSystemsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.FileSystems {
			if *r.OwnerId != client.AccountID {
				continue
			}
			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_fsx_windows_file_system",
				ID:        *r.FileSystemId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
