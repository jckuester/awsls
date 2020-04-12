// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerModel(client *Client) {
	req := client.sagemakerconn.ListModelsRequest(&sagemaker.ListModelsInput{})

	p := sagemaker.NewListModelsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_sagemaker_model:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Models {
			fmt.Println(*r.ModelName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_sagemaker_model: %s", err)
	}

}
