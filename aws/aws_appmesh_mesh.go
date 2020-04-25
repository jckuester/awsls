// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

func ListAppmeshMesh(client *Client) ([]Resource, error) {
	req := client.appmeshconn.ListMeshesRequest(&appmesh.ListMeshesInput{})

	var result []Resource

	p := appmesh.NewListMeshesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Meshes {

			result = append(result, Resource{
				Type:   "aws_appmesh_mesh",
				ID:     *r.MeshName,
				Region: client.appmeshconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
