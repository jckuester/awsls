//+build codegen

package aws

import (
	"reflect"
	"testing"
)

func Test_pluralizeListFunctionCandidateNames(t *testing.T) {
	type args struct {
		terraformType string
	}
	tests := []struct {
		name string
		arg  args
		want []string
	}{
		{
			name: "resource type ending with y",
			arg:  args{terraformType: "aws_gateway"},
			want: []string{"gateways", "gatewayes", "gatewaies"},
		},
		{
			name: "resource type with four underscores",
			arg:  args{terraformType: "aws_storagegateway_nfs_file_share"},
			want: []string{"storagegatewaynfsfileshares", "storagegatewaynfsfilesharees", "nfsfileshares",
				"nfsfilesharees", "fileshares", "filesharees", "shares", "sharees"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pluralizeListFunctionCandidateNames(tt.arg.terraformType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pluralizeListFunctionCandidateNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
