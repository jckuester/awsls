package resource_test

import (
	"reflect"
	"testing"

	"github.com/jckuester/awsls/resource"
	"github.com/stretchr/testify/assert"
)

func TestIsType(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "existing Terraform resource type",
			arg:  "aws_vpc",
			want: true,
		},
		{
			name: "not existing Terraform resource type",
			arg:  "aws_foo",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resource.IsType(tt.arg); got != tt.want {
				t.Errorf("IsType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSupportedType(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "supported resource type",
			arg:  "aws_vpc",
			want: true,
		},
		{
			name: "not supported resource type",
			arg:  "aws_default_vpc",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resource.IsSupportedType(tt.arg); got != tt.want {
				t.Errorf("IsSupportedType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchSupportedTypes(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    []string
		wantErr string
	}{
		{
			name: "no match found",
			arg:  "aws_",
		},
		{
			name:    "invalid glob pattern",
			arg:     "aws_[",
			wantErr: "unexpected end of input",
		},
		{
			name: "single resource matches, no wildcard",
			arg:  "aws_iam_user",
			want: []string{"aws_iam_user"},
		},
		{
			name: "glob pattern with wildcard",
			arg:  "aws_vpc*",
			want: []string{
				"aws_vpc",
				"aws_vpc_endpoint",
				"aws_vpc_endpoint_connection_notification",
				"aws_vpc_endpoint_service",
				"aws_vpc_peering_connection",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resource.MatchSupportedTypes(tt.arg)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchSupportedTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSupportsTags(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "resource type supports tags",
			arg:  "aws_vpc",
			want: true,
		},
		{
			name: "resource type doesn't support tags",
			arg:  "aws_kms_alias",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resource.SupportsTags(tt.arg); got != tt.want {
				t.Errorf("SupportsTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
