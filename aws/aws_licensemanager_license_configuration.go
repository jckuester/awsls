// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
)

func ListLicensemanagerLicenseConfiguration(client *Client) error {
	req := client.licensemanagerconn.ListLicenseConfigurationsRequest(&licensemanager.ListLicenseConfigurationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.LicenseConfigurations) > 0 {
		for _, r := range resp.LicenseConfigurations {
			fmt.Println(*r.LicenseConfigurationArn)
		}
	}

	return nil
}
