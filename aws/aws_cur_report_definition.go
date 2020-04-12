// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

func ListCurReportDefinition(client *Client) error {
	req := client.costandusagereportserviceconn.DescribeReportDefinitionsRequest(&costandusagereportservice.DescribeReportDefinitionsInput{})

	p := costandusagereportservice.NewDescribeReportDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReportDefinitions {
			fmt.Println(*r.ReportName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
