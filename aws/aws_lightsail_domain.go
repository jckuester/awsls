// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailDomain(client *Client) {
	req := client.lightsailconn.GetDomainsRequest(&lightsail.GetDomainsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_lightsail_domain: %s", err)
	} else {
		if len(resp.Domains) > 0 {
			fmt.Println("")
			fmt.Println("aws_lightsail_domain:")
			for _, r := range resp.Domains {
				fmt.Println(*r.Name)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
