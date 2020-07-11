package util_test

import (
	"testing"

	"github.com/jckuester/awsls/util"

	"github.com/jckuester/awsls/aws"
)

func TestNewAWSClientPool(t *testing.T) {
	type args struct {
		profiles []string
		regions  []string
	}
	tests := []struct {
		name    string
		args    args
		want    []aws.Client
		wantErr bool
	}{
		{
			name: "empty list of profiles and regions",
			args: args{
				regions: []string{"us-test-1", "eu-test-2"},
			},
		},
		{
			name: "empty list of regions",
			args: args{
				profiles: []string{"profile1", "profile2"},
			},
		},
		{
			name: "empty list of profiles",
			args: args{
				regions: []string{"us-test-1", "eu-test-2"},
			},
		},
		{
			name: "permutation of multiple profiles and regions",
			args: args{
				profiles: []string{"profile1", "profile2"},
				regions:  []string{"us-test-1", "eu-test-2"},
			},
		},
		{
			name: "permutation of multiple, redundant profiles and regions",
			args: args{
				profiles: []string{"profile1", "profile2", "profile1"},
				regions:  []string{"us-test-1", "eu-test-2", "eu-test-2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := util.NewAWSClientPool(tt.args.profiles, tt.args.regions)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAWSClientPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
