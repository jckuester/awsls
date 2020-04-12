// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmAssociation(client *Client) {
	req := client.ssmconn.ListAssociationsRequest(&ssm.ListAssociationsInput{})

	p := ssm.NewListAssociationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ssm_association:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Associations {
			fmt.Println(*r.AssociationId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ssm_association: %s", err)
	}

}
