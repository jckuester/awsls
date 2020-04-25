// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func ListSesReceiptFilter(client *Client) ([]Resource, error) {
	req := client.sesconn.ListReceiptFiltersRequest(&ses.ListReceiptFiltersInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Filters) > 0 {
		for _, r := range resp.Filters {

			result = append(result, Resource{
				Type:   "aws_ses_receipt_filter",
				ID:     *r.Name,
				Region: client.sesconn.Config.Region,
			})
		}
	}

	return result, nil
}
