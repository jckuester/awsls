// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

func ListQldbLedger(client *Client) {
	req := client.qldbconn.ListLedgersRequest(&qldb.ListLedgersInput{})

	p := qldb.NewListLedgersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_qldb_ledger:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Ledgers {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_qldb_ledger: %s", err)
	}

}
