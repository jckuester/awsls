// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/athena"
)

func ListAthenaWorkgroup(client *Client) error {
	req := client.athenaconn.ListWorkGroupsRequest(&athena.ListWorkGroupsInput{})

	p := athena.NewListWorkGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.WorkGroups {
			fmt.Println(*r.Name)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
