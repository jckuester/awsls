package resource

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/fatih/color"
	awsls "github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/internal"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
	"github.com/jckuester/terradozer/pkg/provider"
)

// UpdatedResources contains resources which Terraform state
// has been updated. Failed updates are logged in the Errors array.
type UpdatedResources struct {
	Resources []awsls.Resource
	Errors    []error
}

// ListInMultipleAccountsAndRegions lists resources of a given resource type in parallel for multiple accounts and
// regions and updates the resources' Terraform state.
func ListInMultipleAccountsAndRegions(rType string, hasAttrs map[string]bool,
	clients map[aws.ClientKey]awsls.Client, providers map[aws.ClientKey]provider.TerraformProvider) UpdatedResources {
	var wg sync.WaitGroup
	sem := internal.NewSemaphore(10)

	result := terraform.ResourcesThreadSafe{
		Resources: []awsls.Resource{},
	}

	for key, client := range clients {
		log.WithFields(log.Fields{
			"type":    rType,
			"region":  key.Region,
			"profile": key.Profile,
			"time":    time.Now().Format("04:05.000"),
		}).Debugf("start listing resources")

		wg.Add(1)

		go func(client awsls.Client) {
			defer wg.Done()

			// Acquire a semaphore so that we can limit concurrency
			sem.Acquire()
			defer sem.Release()

			err := client.SetAccountID()
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))
				return
			}

			resources, err := awsls.ListResourcesByType(&client, rType)
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))
				return
			}

			if len(hasAttrs) > 0 {
				// for performance reasons:
				// only fetch state if some attributes need to be displayed for this resource type
				updatesResources, errs := terraform.UpdateStates(resources, providers, 10)
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
