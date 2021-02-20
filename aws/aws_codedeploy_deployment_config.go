// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodedeployDeploymentConfig(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codedeployconn.ListDeploymentConfigsRequest(&codedeploy.ListDeploymentConfigsInput{})

	var result []terraform.Resource

	p := codedeploy.NewListDeploymentConfigsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DeploymentConfigsList {

			result = append(result, terraform.Resource{
				Type:      "aws_codedeploy_deployment_config",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
