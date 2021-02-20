// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGlueMlTransform(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Glueconn.GetMLTransformsRequest(&glue.GetMLTransformsInput{})

	var result []terraform.Resource

	p := glue.NewGetMLTransformsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Transforms {

			result = append(result, terraform.Resource{
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
