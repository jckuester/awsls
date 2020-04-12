// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func ListSesReceiptFilter(client *Client) {
	req := client.sesconn.ListReceiptFiltersRequest(&ses.ListReceiptFiltersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ses_receipt_filter: %s", err)
	} else {
		if len(resp.Filters) > 0 {
			fmt.Println("")
			fmt.Println("aws_ses_receipt_filter:")
			for _, r := range resp.Filters {
				fmt.Println(*r.Name)

			}
		}
	}

}
