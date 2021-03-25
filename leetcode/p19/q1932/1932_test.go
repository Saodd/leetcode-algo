package q1932

import (
	"github.com/Saodd/leetcode-algo/helper/treenode"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "官方用例1",
			args: args{treenode.Unmarshal([]string{"2", "1", "3"})},
			want: true,
		},
		{
			name: "官方用例2",
			args: args{treenode.Unmarshal([]string{"5", "1", "4", "null", "null", "3", "6"})},
			want: false,
		},
		{
			name: "错误1",
			args: args{treenode.Unmarshal([]string{"10", "5", "15", "null", "null", "6", "20"})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
