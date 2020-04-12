// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func ListEfsFileSystem(client *Client) {
	req := client.efsconn.DescribeFileSystemsRequest(&efs.DescribeFileSystemsInput{})

	p := efs.NewDescribeFileSystemsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_efs_file_system:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.FileSystems {
			fmt.Println(*r.FileSystemId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_efs_file_system: %s", err)
	}

}
