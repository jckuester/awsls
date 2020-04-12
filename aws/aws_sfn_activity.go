// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func ListSfnActivity(client *Client) {
	req := client.sfnconn.ListActivitiesRequest(&sfn.ListActivitiesInput{})

	p := sfn.NewListActivitiesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_sfn_activity:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Activities {
			fmt.Println(*r.ActivityArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_sfn_activity: %s", err)
	}

}
