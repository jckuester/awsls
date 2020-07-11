package util

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/jckuester/awsls/aws"
)

// clientPoolThreadSafe is a concurrent slice implementation to store multiple AWS clients.
type clientPoolThreadSafe struct {
	sync.Mutex
	clients []aws.Client
}

// NewAWSClientPool creates a set of AWS clients for each permutation of the given profiles and regions.
// If profiles, regions, or both are empty, credentials or regions are picked up via the usual default provider chain,
// respectively. For example, if regions are empty, the default region for each profile is used from `~/.aws/config`
// or from the according region environment variable.
func NewAWSClientPool(profiles []string, regions []string) ([]aws.Client, error) {
	errors := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup

	clientPool := &clientPoolThreadSafe{}

	if len(profiles) > 0 && len(regions) > 0 {
		wg.Add(len(profiles) * len(regions))

		for _, profile := range profiles {
			for _, region := range regions {

				go func(p string, r string) {
					defer wg.Done()

					client, err := aws.NewClient(
						external.WithSharedConfigProfile(p),
						external.WithRegion(r))
					if err != nil {
						errors <- err
						return
					}

					clientPool.Lock()
					clientPool.clients = append(clientPool.clients, *client)
					clientPool.Unlock()
				}(profile, region)
			}
		}
	} else if len(profiles) > 0 {
		wg.Add(len(profiles))

		for _, profile := range profiles {
			go func(p string) {
				defer wg.Done()

				client, err := aws.NewClient(external.WithSharedConfigProfile(p))
				if err != nil {
					errors <- err
					return
				}

				clientPool.Lock()
				clientPool.clients = append(clientPool.clients, *client)
				clientPool.Unlock()
			}(profile)
		}
	} else if len(regions) > 0 {
		wg.Add(len(regions))

		for _, region := range regions {
			go func(r string) {
				defer wg.Done()

				client, err := aws.NewClient(external.WithRegion(r))
				if err != nil {
					errors <- err
					return
				}

				clientPool.Lock()
				clientPool.clients = append(clientPool.clients, *client)
				clientPool.Unlock()
			}(region)
		}
	} else {
		client, err := aws.NewClient()
		if err != nil {
			return nil, err
		}

		return []aws.Client{*client}, nil
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-errors:
		close(errors)
		return nil, err
	}

	return clientPool.clients, nil
}
