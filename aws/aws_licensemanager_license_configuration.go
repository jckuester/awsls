// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLicensemanagerLicenseConfiguration(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Licensemanagerconn.ListLicenseConfigurations(ctx, &licensemanager.ListLicenseConfigurationsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.LicenseConfigurations) > 0 {

		for _, r := range resp.LicenseConfigurations {

			result = append(result, terraform.Resource{
				Type:      "aws_licensemanager_license_configuration",
				ID:        *r.LicenseConfigurationArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
