// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueSecurityConfiguration(client *Client) error {
	req := client.glueconn.GetSecurityConfigurationsRequest(&glue.GetSecurityConfigurationsInput{})

	p := glue.NewGetSecurityConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecurityConfigurations {
			fmt.Println(*r.Name)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
