package resource

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/apex/log"
	"github.com/fatih/color"
	awsls "github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/internal"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

// UpdatedResources contains resources which Terraform state
// has been updated. Failed updates are logged in the Errors array.
type UpdatedResources struct {
	Resources []terraform.Resource
	Errors    []error
}

// ListInMultipleAccountsAndRegions lists resources of a given resource type in parallel for multiple accounts and
// regions and updates the resources' Terraform state.
func ListInMultipleAccountsAndRegions(ctx context.Context, rType string, hasAttrs map[string]bool,
	clients map[aws.ClientKey]aws.Client, providerPath string) UpdatedResources {
	var wg sync.WaitGroup
	sem := internal.NewSemaphore(10)

	result := terraform.ResourcesThreadSafe{
		Resources: []terraform.Resource{},
	}

	for key, client := range clients {
		log.WithFields(log.Fields{
			"type":    rType,
			"region":  key.Region,
			"profile": key.Profile,
		}).Debugf("start listing resources")

		wg.Add(1)

		go func(client aws.Client) {
			defer wg.Done()

			// Acquire a semaphore so that we can limit concurrency
			sem.Acquire()
			defer sem.Release()

			err := client.SetAccountID(ctx)
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))
				return
			}

			resources, err := awsls.ListResourcesByType(ctx, &client, rType)
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))
				return
			}

			if len(hasAttrs) > 0 {
				// for performance reasons:
				// only fetch state if some attributes need to be displayed for this resource type
				updatesResources, errs := terraform.UpdateStates(resources, providerPath, 10, true)
				resources = updatesResources

				result.Lock()
				result.Errors = append(result.Errors, errs...)
				result.Unlock()
			}

			result.Lock()
			result.Resources = append(result.Resources, resources...)
			result.Unlock()
		}(client)
	}

	// Wait until listing resources of this type completes for every account and region
	wg.Wait()

	return UpdatedResources{result.Resources, result.Errors}
}
