// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsPlatformApplication(client *Client) error {
	req := client.snsconn.ListPlatformApplicationsRequest(&sns.ListPlatformApplicationsInput{})

	p := sns.NewListPlatformApplicationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.PlatformApplications {
			fmt.Println(*r.PlatformApplicationArn)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
