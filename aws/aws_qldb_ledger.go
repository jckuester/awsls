// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

func ListQldbLedger(client *Client) ([]Resource, error) {
	req := client.qldbconn.ListLedgersRequest(&qldb.ListLedgersInput{})

	var result []Resource

	p := qldb.NewListLedgersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Ledgers {

			result = append(result, Resource{
				Type: "aws_qldb_ledger",
				ID:   *r.Name,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
