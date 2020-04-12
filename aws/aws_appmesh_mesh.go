// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

func ListAppmeshMesh(client *Client) {
	req := client.appmeshconn.ListMeshesRequest(&appmesh.ListMeshesInput{})

	p := appmesh.NewListMeshesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_appmesh_mesh:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Meshes {
			fmt.Println(*r.MeshName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_appmesh_mesh: %s", err)
	}

}
