// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func ListCodecommitRepository(client *Client) {
	req := client.codecommitconn.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})

	p := codecommit.NewListRepositoriesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_codecommit_repository:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Repositories {
			fmt.Println(*r.RepositoryName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_codecommit_repository: %s", err)
	}

}
