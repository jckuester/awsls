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

	testVars := Init(t)

	terraformDir := "./test-fixtures/multiple-profiles-and-regions"

	terraformOptions := GetTerraformOptions(terraformDir, testVars, map[string]interface{}{
		"profile1": testVars.AWSProfile1,
		"profile2": testVars.AWSProfile2,
		"region1":  testVars.AWSRegion1,
		"region2":  testVars.AWSRegion2,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	tests := []struct {
		name            string
		args            []string
		envs            map[string]string
		expectedLogs    []string
		expectedErrCode int
	}{
		{
			name: "multiple profiles and regions via flag",
			args: []string{
				"-p", fmt.Sprintf("%s,%s", testVars.AWSProfile1, testVars.AWSProfile2),
				"-r", fmt.Sprintf("%s,%s", testVars.AWSRegion1, testVars.AWSRegion2),
				"-a", "tags", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile1, testVars.AWSRegion2),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile2, testVars.AWSRegion1),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile2, testVars.AWSRegion2),
			},
		},
		{
			name: "profile via env, multiple regions via flag",
			args: []string{
				"-r", fmt.Sprintf("%s,%s", testVars.AWSRegion1, testVars.AWSRegion2),
				"-a", "tags", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile1, testVars.AWSRegion2),
			},
			envs: map[string]string{
				"AWS_PROFILE": testVars.AWSProfile1,
			},
		},
		{
			name: "profile and region via env",
			args: []string{"-a", "tags", "aws_vpc"},
			envs: map[string]string{
				"AWS_PROFILE":        testVars.AWSProfile1,
				"AWS_DEFAULT_REGION": testVars.AWSRegion2,
			},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile1, testVars.AWSRegion2),
			},
		},
		{
			name: "profile via flag, using default region from AWS config file",
			args: []string{"-a", "tags", "aws_vpc"},
			envs: map[string]string{
				"AWS_PROFILE": testVars.AWSProfile1,
			},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile1, testVars.AWSRegion1),
			},
		},
		{
			name: "all-profiles flag, default region for each profile",
			args: []string{"--all-profiles", "-a", "tags", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("%[1]s\\s+%[2]s\\s+N/A\\s+awsls=test-acc,foo=%[1]s-%[2]s",
					testVars.AWSProfile2, testVars.AWSRegion2),
			},
		},
		{
			name: "with profiles and all-profiles flag",
			args: []string{"-p", "foo", "--all-profiles", "aws_vpc"},
			expectedLogs: []string{
				"--profiles and --all-profiles flag cannot be used together",
			},
			expectedErrCode: 1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := UnsetAWSEnvs()
			require.NoError(t, err)

			err = SetMultiEnvs(tc.envs)
			require.NoError(t, err)

			logBuffer, err := runBinary(t, tc.args...)

			if tc.expectedErrCode > 0 {
				require.EqualError(t, err, "exit status 1")
			} else {
				require.NoError(t, err)
			}

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

	testVars := Init(t)

	terraformDir := "./test-fixtures/tag-attributes"

	terraformOptions := GetTerraformOptions(terraformDir, testVars)

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualVpcIDSingleTag := terraform.Output(t, terraformOptions, "vpc_id_single_tag")
	AssertVpcExists(t, actualVpcIDSingleTag, testVars.AWSProfile1, testVars.AWSRegion1)

	actualVpcIDMultipleTags := terraform.Output(t, terraformOptions, "vpc_id_multiple_tags")
	AssertVpcExists(t, actualVpcIDMultipleTags, testVars.AWSProfile1, testVars.AWSRegion1)

	tests := []struct {
		name         string
		args         []string
		expectedLogs []string
	}{
		{
			name: "without attributes",
			args: []string{
				"-p", testVars.AWSProfile1, "-r", testVars.AWSRegion1, "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A",
					actualVpcIDSingleTag, testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A",
					actualVpcIDMultipleTags, testVars.AWSProfile1, testVars.AWSRegion1),
			},
		},
		{
			name: "string attribute",
			args: []string{
				"-p", testVars.AWSProfile1, "-r", testVars.AWSRegion1,
				"-a", "cidr_block", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+CIDR_BLOCK",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A\\s+10.0.0.0/16",
					actualVpcIDSingleTag, testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A\\s+10.0.0.0/16",
					actualVpcIDSingleTag, testVars.AWSProfile1, testVars.AWSRegion1),
			},
		},
		{
			name: "map attribute",
			args: []string{
				"-p", testVars.AWSProfile1, "-r", testVars.AWSRegion1,
				"--attributes", "tags", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A\\s+foo=bar",
					actualVpcIDSingleTag, testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A\\s+bar=baz,foo=bar",
					actualVpcIDMultipleTags, testVars.AWSProfile1, testVars.AWSRegion1),
			},
		},
		{
			name: "multiple attributes",
			args: []string{
				"-p", testVars.AWSProfile1, "-r", testVars.AWSRegion1,
				"-a", "tags,cidr_block", "aws_vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED\\s+TAGS\\s+CIDR_BLOCK",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A\\s+foo=bar\\s+10.0.0.0/16",
					actualVpcIDSingleTag, testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A\\s+bar=baz,foo=bar\\s+10.0.0.0/16",
					actualVpcIDMultipleTags, testVars.AWSProfile1, testVars.AWSRegion1),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			logBuffer, err := runBinary(t, tc.args...)
			require.NoError(t, err)

			// just to be extra safe: check that nothing is deleted
			AssertVpcExists(t, actualVpcIDSingleTag, testVars.AWSProfile1, testVars.AWSRegion1)
			AssertVpcExists(t, actualVpcIDMultipleTags, testVars.AWSProfile1, testVars.AWSRegion1)

			actualLogs := logBuffer.String()

			for _, expectedLogEntry := range tc.expectedLogs {
				assert.Regexp(t, regexp.MustCompile(expectedLogEntry), actualLogs)
			}

			fmt.Println(actualLogs)
		})
	}
}

func TestAcc_ResourceTypeGlobPattern(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping acceptance test.")
	}

	testVars := Init(t)

	terraformDir := "./test-fixtures/multiple-resource-types"

	terraformOptions := GetTerraformOptions(terraformDir, testVars)

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualVpcID := terraform.Output(t, terraformOptions, "vpc_id")
	actualSubnetID := terraform.Output(t, terraformOptions, "subnet_id")

	tests := []struct {
		name         string
		args         []string
		expectedLogs []string
	}{
		{
			name: "without aws_prefix",
			args: []string{
				"-p", testVars.AWSProfile1, "-r", testVars.AWSRegion1, "vpc"},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A",
					actualVpcID, testVars.AWSProfile1, testVars.AWSRegion1),
			},
		},
		{
			name: "multiple resource types",
			args: []string{
				"-p", testVars.AWSProfile1, "-r", testVars.AWSRegion1, "\"{vpc,subnet}\""},
			expectedLogs: []string{
				"TYPE\\s+ID\\s+PROFILE\\s+REGION\\s+CREATED",
				fmt.Sprintf("aws_vpc\\s+%s\\s+%s\\s+%s\\s+N/A",
					actualVpcID, testVars.AWSProfile1, testVars.AWSRegion1),
				fmt.Sprintf("aws_subnet\\s+%s\\s+%s\\s+%s\\s+N/A",
					actualSubnetID, testVars.AWSProfile1, testVars.AWSRegion1),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			logBuffer, err := runBinary(t, tc.args...)
			require.NoError(t, err)

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
