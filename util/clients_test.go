package util_test

import (
	"testing"

	"github.com/jckuester/awsls/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAWSClientPool(t *testing.T) {
	type args struct {
		profiles []string
		regions  []string
	}
	tests := []struct {
		name    string
		args    args
		envs    map[string]string
		want    []util.AWSClientKey
		wantErr bool
	}{
		{
			name: "no profiles and regions via flag, default region via env",
			args: args{},
			envs: map[string]string{
				"AWS_DEFAULT_REGION": "us-test-1",
			},
			want: []util.AWSClientKey{
				{"", "us-test-1"},
			},
		},
		{
			name: "profiles via flag, default region via config file",
			args: args{
				profiles: []string{"profile1", "profile2"},
			},
			envs: map[string]string{
				"AWS_CONFIG_FILE": "../test/test-fixtures/aws-config",
			},
			want: []util.AWSClientKey{
				{"profile1", "us-test-1"},
				{"profile2", "us-test-2"},
			},
		},
		{
			name: "profiles via flag, default region via env",
			args: args{
				profiles: []string{"profile1", "profile2"},
			},
			envs: map[string]string{
				"AWS_DEFAULT_REGION": "us-test-3",
				"AWS_CONFIG_FILE":    "../test/test-fixtures/aws-config",
			},
			want: []util.AWSClientKey{
				{"profile1", "us-test-3"},
				{"profile2", "us-test-3"},
			},
		},
		{
			name: "no profiles and regions via flag, profile via env",
			args: args{},
			envs: map[string]string{
				"AWS_CONFIG_FILE": "../test/test-fixtures/aws-config",
				"AWS_PROFILE":     "profile1",
			},
			// Note: unfortunately, if the profile is not explicitly added to the config
			// we cannot retrieve the profile name from the config ex-post
			want: []util.AWSClientKey{
				{"", "us-test-1"},
			},
		},
		{
			name: "no profiles but regions via flag",
			args: args{
				regions: []string{"us-test-1", "us-test-2"},
			},
			want: []util.AWSClientKey{
				{"", "us-test-1"},
				{"", "us-test-2"},
			},
		},
		{
			name: "permutation of multiple profiles and regions via flag",
			args: args{
				profiles: []string{"profile1", "profile2"},
				regions:  []string{"us-test-1", "us-test-2"},
			},
			want: []util.AWSClientKey{
				{"profile1", "us-test-1"},
				{"profile1", "us-test-2"},
				{"profile2", "us-test-1"},
				{"profile2", "us-test-2"},
			},
		},
		{
			name: "permutation of multiple, duplicate profiles and regions via flag",
			args: args{
				profiles: []string{"profile1", "profile2", "profile1"},
				regions:  []string{"us-test-1", "us-test-2", "us-test-2"},
			},
			want: []util.AWSClientKey{
				{"profile1", "us-test-1"},
				{"profile1", "us-test-2"},
				{"profile2", "us-test-1"},
				{"profile2", "us-test-2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UnsetMultiEnvs([]string{
				"AWS_DEFAULT_REGION",
				"AWS_REGION",
				"AWS_PROFILE",
				"AWS_SDK_LOAD_CONFIG",
				"AWS_ACCESS_KEY_ID",
				"AWS_SECRET_ACCESS_KEY",
				"AWS_SESSION_TOKEN",
			})
			require.NoError(t, err)

			err = SetMultiEnvs(tt.envs)
			require.NoError(t, err)

			got, err := util.NewAWSClientPool(tt.args.profiles, tt.args.regions)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAWSClientPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Len(t, got, len(tt.want))

			for _, clientKey := range tt.want {
				actualClient, ok := got[clientKey]
				if !ok {
					t.Fatal("AWS client does not exist")
				}
				assert.Equal(t, clientKey.Profile, actualClient.Profile)
				assert.Equal(t, clientKey.Region, actualClient.Region)
			}
		})
	}
}
