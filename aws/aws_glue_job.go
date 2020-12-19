// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueJob(client *Client) ([]Resource, error) {
	req := client.Glueconn.GetJobsRequest(&glue.GetJobsInput{})

	var result []Resource

	p := glue.NewGetJobsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Jobs {

			result = append(result, Resource{
				Type:      "aws_glue_job",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
