// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

func ListQldbLedger(client *Client) error {
	req := client.qldbconn.ListLedgersRequest(&qldb.ListLedgersInput{})

	p := qldb.NewListLedgersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Ledgers {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
