// +build codegen

package aws

import "testing"

func TestListFunctionName(t *testing.T) {
	tests := []struct {
		name  string
		rType ResourceType
		want  string
	}{
		{
			name: "aws_iam_role",
			rType: ResourceType{
				Name: "aws_iam_role",
			},
			want: "ListIamRole",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rType.ListFunctionName(); got != tt.want {
				t.Errorf("ListFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
