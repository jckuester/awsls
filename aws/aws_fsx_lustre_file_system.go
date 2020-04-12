// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

func ListFsxLustreFileSystem(client *Client) error {
	req := client.fsxconn.DescribeFileSystemsRequest(&fsx.DescribeFileSystemsInput{})

	p := fsx.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.FileSystems {
			fmt.Println(*r.FileSystemId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.CreationTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
