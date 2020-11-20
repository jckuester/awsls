// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueMlTransform(client *Client) ([]Resource, error) {
	req := client.Glueconn.GetMLTransformsRequest(&glue.GetMLTransformsInput{})

	var result []Resource

	p := glue.NewGetMLTransformsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Transforms {

			result = append(result, Resource{
				Type:      "aws_glue_ml_transform",
				ID:        *r.TransformId,
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
