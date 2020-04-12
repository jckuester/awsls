// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
)

func ListLicensemanagerLicenseConfiguration(client *Client) {
	req := client.licensemanagerconn.ListLicenseConfigurationsRequest(&licensemanager.ListLicenseConfigurationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_licensemanager_license_configuration: %s", err)
	} else {
		if len(resp.LicenseConfigurations) > 0 {
			fmt.Println("")
			fmt.Println("aws_licensemanager_license_configuration:")
			for _, r := range resp.LicenseConfigurations {
				fmt.Println(*r.LicenseConfigurationArn)

			}
		}
	}

}
