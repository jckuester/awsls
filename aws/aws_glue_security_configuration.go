// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueSecurityConfiguration(client *Client) {
	req := client.glueconn.GetSecurityConfigurationsRequest(&glue.GetSecurityConfigurationsInput{})

	p := glue.NewGetSecurityConfigurationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_glue_security_configuration:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecurityConfigurations {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_glue_security_configuration: %s", err)
	}

}
