// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerModel(client *Client) ([]Resource, error) {
	req := client.Sagemakerconn.ListModelsRequest(&sagemaker.ListModelsInput{})

	var result []Resource

	p := sagemaker.NewListModelsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Models {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_sagemaker_model",
				ID:        *r.ModelName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
