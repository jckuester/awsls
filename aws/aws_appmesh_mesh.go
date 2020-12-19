// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

func ListAppmeshMesh(client *Client) ([]Resource, error) {
	req := client.Appmeshconn.ListMeshesRequest(&appmesh.ListMeshesInput{})

	var result []Resource

	p := appmesh.NewListMeshesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Meshes {

			result = append(result, Resource{
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
