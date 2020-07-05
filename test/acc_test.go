package test

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/onsi/gomega/gexec"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	packagePath = "github.com/jckuester/awsls"
)

func TestAcc_Attributes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping acceptance test.")
	}

	env := InitEnv(t)

	terraformDir := "./test-fixtures/aws-vpc"

	terraformOptions := GetTerraformOptions(terraformDir, env)

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualVpcID := terraform.Output(t, terraformOptions, "vpc_id")
	AssertVpcExists(t, actualVpcID, env)

	tests := []struct {
		name         string
		attributes   []string
		expectedLogs []string
	}{
		{
			name:       "single attributes",
			attributes: []string{"--attributes", "tags", "aws_vpc"},
			expectedLogs: []string{
				"REGION      CREATED   TAGS",
				"us-west-2   N/A       bar=baz,foo=bar",
			},
		},
		{
			name:       "multiple attributes",
			attributes: []string{"-a", "tags,cidr_block", "aws_vpc"},
			expectedLogs: []string{
				"REGION      CREATED   TAGS              CIDR_BLOCK",
				"us-west-2   N/A       bar=baz,foo=bar   10.0.0.0/16",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			logBuffer, err := runBinary(t, tc.attributes...)
			require.NoError(t, err)

			AssertVpcExists(t, actualVpcID, env)

			actualLogs := logBuffer.String()

			for _, expectedLogEntry := range tc.expectedLogs {
				assert.Contains(t, actualLogs, expectedLogEntry)
			}

			fmt.Println(actualLogs)
		})
	}
}

func TestAcc_NonExistingResourceType(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping acceptance test.")
	}

	logBuffer, err := runBinary(t, "aws_foo")
	require.Error(t, err)

	actualLogs := logBuffer.String()

	assert.Contains(t, actualLogs, "Error: no resource type found: aws_foo")

	fmt.Println(actualLogs)
}

func TestAcc_UnsupportedResourceType(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping acceptance test.")
	}

	logBuffer, err := runBinary(t, "--debug", "aws_opsworks_mysql_layer")
	require.Error(t, err)

	actualLogs := logBuffer.String()

	assert.Contains(t, actualLogs, "resource type not (yet) supported: aws_opsworks_mysql_layer")
	assert.Contains(t, actualLogs, "Error: no resource type found: aws_opsworks_mysql_layer")

	fmt.Println(actualLogs)
}

func TestAcc_Version(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping acceptance test.")
	}

	logBuffer, err := runBinary(t, "--version")
	require.NoError(t, err)

	actualLogs := logBuffer.String()

	assert.Contains(t, actualLogs, fmt.Sprintf(`
version: dev
commit: ?
built at: ?
using: %s`, runtime.Version()))

	fmt.Println(actualLogs)
}

func runBinary(t *testing.T, args ...string) (*bytes.Buffer, error) {
	defer gexec.CleanupBuildArtifacts()

	compiledPath, err := gexec.Build(packagePath)
	require.NoError(t, err)

	logBuffer := &bytes.Buffer{}

	p := exec.Command(compiledPath, args...)
	p.Stdout = logBuffer
	p.Stderr = logBuffer

	err = p.Run()

	return logBuffer, err
}
