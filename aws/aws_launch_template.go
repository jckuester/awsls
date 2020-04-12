// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListLaunchTemplate(client *Client) error {
	req := client.ec2conn.DescribeLaunchTemplatesRequest(&ec2.DescribeLaunchTemplatesInput{})

	p := ec2.NewDescribeLaunchTemplatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LaunchTemplates {
			fmt.Println(*r.LaunchTemplateId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.CreateTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
