// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/emr"
)

func ListEmrSecurityConfiguration(client *Client) error {
	req := client.emrconn.ListSecurityConfigurationsRequest(&emr.ListSecurityConfigurationsInput{})

	p := emr.NewListSecurityConfigurationsPaginator(req)
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
