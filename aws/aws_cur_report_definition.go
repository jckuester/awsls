// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

func ListCurReportDefinition(client *Client) {
	req := client.costandusagereportserviceconn.DescribeReportDefinitionsRequest(&costandusagereportservice.DescribeReportDefinitionsInput{})

	p := costandusagereportservice.NewDescribeReportDefinitionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_cur_report_definition:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReportDefinitions {
			fmt.Println(*r.ReportName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_cur_report_definition: %s", err)
	}

}
