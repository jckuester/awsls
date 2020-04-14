package resource_test

import (
	"testing"

	"github.com/jckuester/awsls/resource"
)

func TestIsResourceType(t *testing.T) {
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
			if got := resource.IsResourceType(tt.givenArg); got != tt.want {
				t.Errorf("IsResourceType() = %v, want %v", got, tt.want)
			}
		})
	}
}
