package util

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/jckuester/terradozer/pkg/provider"
	"github.com/zclconf/go-cty/cty"
)

// providerPoolThreadSafe is a concurrent map implementation to store multiple Terraform AWS Providers.
type providerPoolThreadSafe struct {
	sync.Mutex
	providers map[AWSClientKey]provider.TerraformProvider
}

func NewProviderPool(clientKeys []AWSClientKey) (map[AWSClientKey]provider.TerraformProvider, error) {
	metaPlugin, err := provider.Install("aws", "2.68.0", "~/.awsls", true)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("failed to install provider (%s): %s", "aws", err))
	}

	errors := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup

	providerPool := &providerPoolThreadSafe{
		providers: make(map[AWSClientKey]provider.TerraformProvider),
	}

	if len(clientKeys) > 0 {
		wg.Add(len(clientKeys))

		for _, clientKey := range clientKeys {
			go func(p string, r string) {
				defer wg.Done()

				pr, err := provider.Launch(metaPlugin.Path, 10*time.Second)
				if err != nil {
					errors <- fmt.Errorf("failed to launch provider (%s): %s", metaPlugin.Path, err)
					return
				}

				config := cty.ObjectVal(map[string]cty.Value{
					"profile":                     cty.StringVal(p),
					"region":                      cty.StringVal(r),
					"access_key":                  cty.UnknownVal(cty.DynamicPseudoType),
					"allowed_account_ids":         cty.UnknownVal(cty.DynamicPseudoType),
					"assume_role":                 cty.UnknownVal(cty.DynamicPseudoType),
					"endpoints":                   cty.UnknownVal(cty.DynamicPseudoType),
					"forbidden_account_ids":       cty.UnknownVal(cty.DynamicPseudoType),
					"insecure":                    cty.UnknownVal(cty.DynamicPseudoType),
					"max_retries":                 cty.UnknownVal(cty.DynamicPseudoType),
					"s3_force_path_style":         cty.UnknownVal(cty.DynamicPseudoType),
					"secret_key":                  cty.UnknownVal(cty.DynamicPseudoType),
					"shared_credentials_file":     cty.UnknownVal(cty.DynamicPseudoType),
					"skip_credentials_validation": cty.UnknownVal(cty.DynamicPseudoType),
					"skip_get_ec2_platforms":      cty.UnknownVal(cty.DynamicPseudoType),
					"skip_metadata_api_check":     cty.UnknownVal(cty.DynamicPseudoType),
					"skip_region_validation":      cty.UnknownVal(cty.DynamicPseudoType),
					"skip_requesting_account_id":  cty.UnknownVal(cty.DynamicPseudoType),
					"token":                       cty.UnknownVal(cty.DynamicPseudoType),
					"ignore_tag_prefixes":         cty.UnknownVal(cty.DynamicPseudoType),
					"ignore_tags":                 cty.UnknownVal(cty.DynamicPseudoType),
				})

				err = pr.Configure(config)
				if err != nil {
					errors <- fmt.Errorf("failed to configure provider (name=%s, version=%s): %s",
						metaPlugin.Name, metaPlugin.Version, err)
					return
				}

				providerPool.Lock()
				providerPool.providers[AWSClientKey{p, r}] = *pr
				providerPool.Unlock()
			}(clientKey.Profile, clientKey.Region)
		}
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

	return providerPool.providers, nil
}
