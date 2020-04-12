// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func ListSesConfigurationSet(client *Client) {
	req := client.sesconn.ListConfigurationSetsRequest(&ses.ListConfigurationSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ses_configuration_set: %s", err)
	} else {
		if len(resp.ConfigurationSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_ses_configuration_set:")
			for _, r := range resp.ConfigurationSets {
				fmt.Println(*r.Name)

			}
		}
	}

}
