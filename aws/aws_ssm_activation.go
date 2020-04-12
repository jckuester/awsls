// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmActivation(client *Client) {
	req := client.ssmconn.DescribeActivationsRequest(&ssm.DescribeActivationsInput{})

	p := ssm.NewDescribeActivationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ssm_activation:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ActivationList {
			fmt.Println(*r.ActivationId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ssm_activation: %s", err)
	}

}
