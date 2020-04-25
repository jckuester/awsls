package resource_test

import (
	"reflect"
	"testing"

	"github.com/jckuester/awsls/resource"
)

func TestIsType(t *testing.T) {
	tests := []struct {
		name     string
		givenArg string
		want     bool
	}{
		{
			name:     "valid resource type",
			givenArg: "aws_vpc",
			want:     true,
		},
		{
			name:     "non-existing resource type",
			givenArg: "aws_foo",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resource.IsType(tt.givenArg); got != tt.want {
				t.Errorf("IsResourceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchTypes(t *testing.T) {
	tests := []struct {
		name     string
		givenArg string
		want     []string
	}{
		{
			name:     "no matches found",
			givenArg: "aws_",
			want:     []string{},
		},
		{
			name:     "with wildcard",
			givenArg: "aws_iam_user*",
			want: []string{
				"aws_iam_user_group_membership",
				"aws_iam_user_policy_attachment",
				"aws_iam_user_policy",
				"aws_iam_user_ssh_key",
				"aws_iam_user",
				"aws_iam_user_login_profile",
			},
		},
		{
			name:     "single resource, without wildcard",
			givenArg: "aws_iam_user",
			want:     []string{"aws_iam_user"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resource.MatchTypes(tt.givenArg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
