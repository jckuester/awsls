// Package test contains acceptance tests.
package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertVpcExists(t *testing.T, actualVpcID, profile, region string) {
	err := os.Setenv("AWS_PROFILE", profile)
	require.NoError(t, err)

	_, err = aws.GetVpcByIdE(t, actualVpcID, region)
	assert.NoError(t, err, "resource has been unexpectedly deleted")

	err = os.Unsetenv("AWS_PROFILE")
	require.NoError(t, err)
}
