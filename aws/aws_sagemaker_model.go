// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerModel(client *Client) error {
	req := client.sagemakerconn.ListModelsRequest(&sagemaker.ListModelsInput{})

	p := sagemaker.NewListModelsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Models {
			fmt.Println(*r.ModelName)

			fmt.Printf("CreatedAt: %s\n", *r.CreationTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
