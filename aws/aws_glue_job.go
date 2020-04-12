// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueJob(client *Client) {
	req := client.glueconn.GetJobsRequest(&glue.GetJobsInput{})

	p := glue.NewGetJobsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_glue_job:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Jobs {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_glue_job: %s", err)
	}

}
