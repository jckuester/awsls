// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListQldbLedger(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Qldbconn.ListLedgersRequest(&qldb.ListLedgersInput{})

	var result []terraform.Resource

	p := qldb.NewListLedgersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Ledgers {

			result = append(result, terraform.Resource{
				Type:      "aws_qldb_ledger",
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
