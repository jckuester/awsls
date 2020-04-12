// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

func ListAppmeshMesh(client *Client) error {
	req := client.appmeshconn.ListMeshesRequest(&appmesh.ListMeshesInput{})

	p := appmesh.NewListMeshesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Meshes {
			fmt.Println(*r.MeshName)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
