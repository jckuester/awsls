// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLexIntent(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Lexmodelbuildingserviceconn.GetIntentsRequest(&lexmodelbuildingservice.GetIntentsInput{})

	var result []terraform.Resource

	p := lexmodelbuildingservice.NewGetIntentsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Intents {

			result = append(result, terraform.Resource{
				Type:      "aws_lex_intent",
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
