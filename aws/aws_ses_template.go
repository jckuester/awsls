// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func ListSesTemplate(client *Client) {
	req := client.sesconn.ListTemplatesRequest(&ses.ListTemplatesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ses_template: %s", err)
	} else {
		if len(resp.TemplatesMetadata) > 0 {
			fmt.Println("")
			fmt.Println("aws_ses_template:")
			for _, r := range resp.TemplatesMetadata {
				fmt.Println(*r.Name)

			}
		}
	}

}
