// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func ListSesReceiptFilter(client *Client) error {
	req := client.sesconn.ListReceiptFiltersRequest(&ses.ListReceiptFiltersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Filters) > 0 {
		for _, r := range resp.Filters {
			fmt.Println(*r.Name)

		}
	}

	return nil
}
