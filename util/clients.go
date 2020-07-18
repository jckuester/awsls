package util

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/jckuester/awsls/aws"
)

// AWSClientPoolThreadSafe is a concurrent map implementation to store multiple AWS clients.
type AWSClientPoolThreadSafe struct {
	sync.Mutex
	clients map[AWSClientKey]aws.Client
}

type AWSClientKey struct {
	Profile, Region string
}

// NewAWSClientPool creates an AWS client for each permutation of the given profiles and regions.
// If profiles, regions, or both are empty, credentials and regions are picked up via the usual default provider chain,
// respectively. For example, if regions are empty, the default region for each profile is used from `~/.aws/config`
// or from the according region environment variable.
func NewAWSClientPool(profiles []string, regions []string) (map[AWSClientKey]aws.Client, error) {
	errors := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup

	clientPool := &AWSClientPoolThreadSafe{
		clients: make(map[AWSClientKey]aws.Client),
	}

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
					clientPool.clients[AWSClientKey{p, client.Region}] = *client
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
				clientPool.clients[AWSClientKey{p, client.Region}] = *client
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
				clientPool.clients[AWSClientKey{"", client.Region}] = *client
				clientPool.Unlock()
			}(region)
		}
	} else {
		client, err := aws.NewClient()
		if err != nil {
			return nil, err
		}

		return map[AWSClientKey]aws.Client{AWSClientKey{"", client.Region}: *client}, nil
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
