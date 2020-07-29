package test

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
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

func TestAcc_ProfilesAndRegions(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping acceptance test.")
	}

	env := InitEnv(t)

	terraformDir := "./test-fixtures/multiple-profiles-and-regions"

	terraformOptions := GetTerraformOptions(terraformDir, env, map[string]interface{}{
		"profile1": profile1,
		"region1":  region1,
		"profile2": profile2,
		"region2":  region2,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	tests := []struct {
		name         string
		args         []string
		envs         map[string]string
		expectedLogs []string
	}{
		{
			name: "multiple profiles and regions via flag",
			args: []string{
				"-p", fmt.Sprintf("%s,%s", profile1, profile2),
				"-r", fmt.Sprintf("%s,%s", region1, region2),
				"-a", "tags", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile1, region1),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile1, region2),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile2, region1),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile2, region2),
			},
		},
		{
			name: "profile via env, multiple regions via flag",
			args: []string{
				"-r", fmt.Sprintf("%s,%s", region1, region2),
				"-a", "tags", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile1, region1),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile1, region2),
			},
			envs: map[string]string{
				"AWS_PROFILE": profile1,
			},
		},
		{
			name: "profile and region via env",
			args: []string{"-a", "tags", "aws_vpc"},
			envs: map[string]string{
				"AWS_PROFILE":        profile1,
				"AWS_DEFAULT_REGION": region2,
			},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile1, region2),
			},
		},
		{
			name: "profile via flag, using default region from AWS config file",
			args: []string{"-a", "tags", "aws_vpc"},
			envs: map[string]string{
				"AWS_PROFILE": profile1,
			},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s", profile1, region1),
			},
		},
		{
			name: "all-profiles flag, default region for each profile",
			args: []string{"--all-profiles", "aws_vpc"},
			envs: map[string]string{
				"AWS_CONFIG_FILE": "../test/test-fixtures/aws-config",
			},
			expectedLogs: []string{},
		},
		{
			name: "with profiles and all-profiles flag",
			args: []string{"-p", "foo", "--all-profiles", "aws_vpc"},
			expectedLogs: []string{
				"--profiles and --all-profiles flag cannot be used together",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := UnsetAWSEnvs()
			require.NoError(t, err)

			err = SetMultiEnvs(tc.envs)
			require.NoError(t, err)

			logBuffer, err := runBinary(t, tc.args...)
			require.NoError(t, err)

			actualLogs := logBuffer.String()

			for _, expectedLogEntry := range tc.expectedLogs {
				assert.Regexp(t, regexp.MustCompile(expectedLogEntry), actualLogs)
			}

			fmt.Println(actualLogs)

			err = UnsetAWSEnvs()
			require.NoError(t, err)
		})
	}
}

func TestAcc_Attributes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping acceptance test.")
	}

	env := InitEnv(t)

	terraformDir := "./test-fixtures/tag-attributes"

	terraformOptions := GetTerraformOptions(terraformDir, env)

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualVpcID1 := terraform.Output(t, terraformOptions, "vpc_id1")
	AssertVpcExists(t, actualVpcID1, env)

	actualVpcID2 := terraform.Output(t, terraformOptions, "vpc_id2")
	AssertVpcExists(t, actualVpcID2, env)

	tests := []struct {
		name         string
		args         []string
		expectedLogs []string
	}{
		{
			name: "without attributes",
			args: []string{
				"-p", env.AWSProfile, "-r", env.AWSRegion, "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A", actualVpcID1, env.AWSRegion),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A", actualVpcID2, env.AWSRegion),
			},
		},
		{
			name: "string attribute",
			args: []string{
				"-p", env.AWSProfile, "-r", env.AWSRegion,
				"-a", "cidr_block", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+CIDR_BLOCK",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A\\s+10.0.0.0/16", actualVpcID1, env.AWSRegion),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A\\s+10.0.0.0/16", actualVpcID1, env.AWSRegion),
			},
		},
		{
			name: "map attribute",
			args: []string{"--attributes", "tags", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A\\s+foo=bar", actualVpcID1, env.AWSRegion),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A\\s+bar=baz,foo=bar", actualVpcID1, env.AWSRegion),
			},
		},
		{
			name: "multiple attributes",
			args: []string{"-a", "tags,cidr_block", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS\\s+CIDR_BLOCK",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A\\sfoo=bar\\s+10.0.0.0/16", actualVpcID1, env.AWSRegion),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+N/A\\s+bar=baz,foo=bar\\s+10.0.0.0/16", actualVpcID1, env.AWSRegion),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			logBuffer, err := runBinary(t, tc.args...)
			require.NoError(t, err)

			// just to be extra safe: check that nothing is deleted
			AssertVpcExists(t, actualVpcID1, env)
			AssertVpcExists(t, actualVpcID2, env)

			actualLogs := logBuffer.String()

			for _, expectedLogEntry := range tc.expectedLogs {
				assert.Regexp(t, regexp.MustCompile(expectedLogEntry), actualLogs)
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
