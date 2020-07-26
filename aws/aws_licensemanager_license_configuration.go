// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
)

func ListLicensemanagerLicenseConfiguration(client *Client) ([]Resource, error) {
	req := client.Licensemanagerconn.ListLicenseConfigurationsRequest(&licensemanager.ListLicenseConfigurationsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.LicenseConfigurations) > 0 {
		for _, r := range resp.LicenseConfigurations {

			result = append(result, Resource{
				Type:    "aws_licensemanager_license_configuration",
				ID:      *r.LicenseConfigurationArn,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	return result, nil
}
