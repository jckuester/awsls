// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/emr"
)

func ListEmrSecurityConfiguration(client *Client) {
	req := client.emrconn.ListSecurityConfigurationsRequest(&emr.ListSecurityConfigurationsInput{})

	p := emr.NewListSecurityConfigurationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_emr_security_configuration:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecurityConfigurations {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_emr_security_configuration: %s", err)
	}

}
