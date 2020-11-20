// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
)

func ListLexSlotType(client *Client) ([]Resource, error) {
	req := client.Lexmodelbuildingserviceconn.GetSlotTypesRequest(&lexmodelbuildingservice.GetSlotTypesInput{})

	var result []Resource

	p := lexmodelbuildingservice.NewGetSlotTypesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SlotTypes {

			result = append(result, Resource{
				Type:      "aws_lex_slot_type",
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
