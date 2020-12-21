// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
)

func ListLexIntent(client *Client) ([]Resource, error) {
	req := client.Lexmodelbuildingserviceconn.GetIntentsRequest(&lexmodelbuildingservice.GetIntentsInput{})

	var result []Resource

	p := lexmodelbuildingservice.NewGetIntentsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Intents {

			result = append(result, Resource{
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
