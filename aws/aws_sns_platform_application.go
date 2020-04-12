// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsPlatformApplication(client *Client) {
	req := client.snsconn.ListPlatformApplicationsRequest(&sns.ListPlatformApplicationsInput{})

	p := sns.NewListPlatformApplicationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_sns_platform_application:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.PlatformApplications {
			fmt.Println(*r.PlatformApplicationArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_sns_platform_application: %s", err)
	}

}
