// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func ListEcrRepository(client *Client) {
	req := client.ecrconn.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{})

	p := ecr.NewDescribeRepositoriesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ecr_repository:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Repositories {
			fmt.Println(*r.RepositoryName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ecr_repository: %s", err)
	}

}
