// Package test contains acceptance tests.
package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	// tfstateBucket is the S3 bucket that stores all Terraform states of acceptance tests.
	// Note: the bucket must be located in `profile1` and `region1`.
	tfstateBucket = "awsls-testacc-tfstate-492043"
	// profile1 is used as profile for the 1st test account if not overwritten by TEST_AWS_PROFILE1.
	profile1 = "myaccount1"
	// region1 is used as the 1st test region if not overwritten by TEST_AWS_REGION1.
	region1 = "us-west-2"
	// profile2 is used as profile for the 2nd test account if not overwritten by TEST_AWS_PROFILE2.
	profile2 = "myaccount2"
	// region1 is used as the 2nd test region if not overwritten by TEST_AWS_REGION2.
	region2 = "us-east-1"
)

// Vars are environment variables used in tests.
type Vars struct {
	AWSProfile1 string
	AWSProfile2 string
	AWSRegion1  string
	AWSRegion2  string
}

// Init sets variables for acceptance tests.
// Variables can be overwritten by environment variables.
func Init(t *testing.T) Vars {
	t.Helper()

	return Vars{
		AWSProfile1: getEnvOrDefault(t, "TEST_AWS_PROFILE1", profile1),
		AWSProfile2: getEnvOrDefault(t, "TEST_AWS_PROFILE2", profile2),
		AWSRegion1:  getEnvOrDefault(t, "TEST_AWS_REGION1", region1),
		AWSRegion2:  getEnvOrDefault(t, "TEST_AWS_REGION2", region2),
	}
}

func getEnvOrDefault(t *testing.T, envName, defaultValue string) string {
	varValue := os.Getenv(envName)
	if varValue == "" {
		varValue = defaultValue

		t.Logf("env %s not set, therefore using the following default value: %s",
			envName, defaultValue)
	}
	return varValue
}

func GetTerraformOptions(terraformDir string, env Vars, overrideVars ...map[string]interface{}) *terraform.Options {
	name := "awsls-testacc-" + strings.ToLower(random.UniqueId())

	vars := map[string]interface{}{
		"region":  env.AWSRegion1,
		"profile": env.AWSProfile2,
		"name":    name,
	}

	if len(overrideVars) > 0 {
		vars = overrideVars[0]
	}

	return &terraform.Options{
		TerraformDir: terraformDir,
		NoColor:      true,
		Vars:         vars,
		// BackendConfig defines where to store the Terraform state files of tests
		BackendConfig: map[string]interface{}{
			"bucket":  tfstateBucket,
			"key":     fmt.Sprintf("%s.tfstate", name),
			"region":  env.AWSRegion1,
			"profile": env.AWSProfile2,
			"encrypt": true,
		},
	}
}

func AssertVpcExists(t *testing.T, actualVpcID, profile, region string) {
	err := os.Setenv("AWS_PROFILE", profile)
	require.NoError(t, err)

	_, err = aws.GetVpcByIdE(t, actualVpcID, region)
	assert.NoError(t, err, "resource has been unexpectedly deleted")

	err = os.Unsetenv("AWS_PROFILE")
	require.NoError(t, err)
}

func SetMultiEnvs(variables map[string]string) error {
	for envName, envValue := range variables {
		err := os.Setenv(envName, envValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func UnsetAWSEnvs() error {
	return UnsetMultiEnvs([]string{
		"AWS_DEFAULT_REGION",
		"AWS_REGION",
		"AWS_PROFILE",
		"AWS_SDK_LOAD_CONFIG",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
	})
}

func UnsetMultiEnvs(variables []string) error {
	for _, envName := range variables {
		err := os.Unsetenv(envName)
		if err != nil {
			return err
		}
	}

	return nil
}
