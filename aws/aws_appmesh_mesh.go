// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAppmeshMesh(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Appmeshconn.ListMeshesRequest(&appmesh.ListMeshesInput{})

	var result []terraform.Resource

	p := appmesh.NewListMeshesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Meshes {

			result = append(result, terraform.Resource{
				Type:      "aws_appmesh_mesh",
				ID:        *r.MeshName,
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
