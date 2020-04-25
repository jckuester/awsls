// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerModel(client *Client) ([]Resource, error) {
	req := client.sagemakerconn.ListModelsRequest(&sagemaker.ListModelsInput{})

	var result []Resource

	p := sagemaker.NewListModelsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Models {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:   "aws_sagemaker_model",
				ID:     *r.ModelName,
				Region: client.sagemakerconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
