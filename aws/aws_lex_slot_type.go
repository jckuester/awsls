// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLexSlotType(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Lexmodelbuildingserviceconn.GetSlotTypesRequest(&lexmodelbuildingservice.GetSlotTypesInput{})

	var result []terraform.Resource

	p := lexmodelbuildingservice.NewGetSlotTypesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.SlotTypes {

			result = append(result, terraform.Resource{
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
