package q1929

import (
	"github.com/Saodd/leetcode-algo/helper/treenode"
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "官方用例1",
			args: args{[]int{-10, -3, 0, 5, 9}},
			want: treenode.Unmarshal([]string{"0", "-3", "9", "-10", "null", "5"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", treenode.Marshal(got), treenode.Marshal(tt.want))
			}
		})
	}
}
