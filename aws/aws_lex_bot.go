// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
)

func ListLexBot(client *Client) ([]Resource, error) {
	req := client.Lexmodelbuildingserviceconn.GetBotsRequest(&lexmodelbuildingservice.GetBotsInput{})

	var result []Resource

	p := lexmodelbuildingservice.NewGetBotsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Bots {

			result = append(result, Resource{
				Type:      "aws_lex_bot",
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
