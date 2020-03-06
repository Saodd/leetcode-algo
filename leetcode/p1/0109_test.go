package p1

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "示例1",
			args: args{common.CreateListChain([]int{-10, -3, 0, 5, 9})},
			want: &common.TreeNode{
				Val:   0,
				Left:  &common.TreeNode{Val: -10, Right: &common.TreeNode{Val: -3}},
				Right: &common.TreeNode{Val: 5, Right: &common.TreeNode{Val: 9}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
				common.PrintTreeNodes(got)
				common.PrintTreeNodes(tt.want)
			}
		})
	}
}
