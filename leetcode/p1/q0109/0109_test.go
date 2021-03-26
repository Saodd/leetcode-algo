package q0109

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
	"github.com/Saodd/leetcode-algo/helper/treenode"
	"reflect"
	"testing"
)

func Test_sortedListToBST2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "示例1",
			args: args{listnode.Unmarshal([]int{-10, -3, 0, 5, 9})},
			want: treenode.Unmarshal([]string{"0", "-3", "9", "-10", "null", "5"}), // 一个可能的答案
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", treenode.Marshal(got), treenode.Marshal(tt.want))
			}
		})
	}
}

func Test_sortedListToBST1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "示例1",
			args: args{listnode.Unmarshal([]int{-10, -3, 0, 5, 9})},
			want: treenode.Unmarshal([]string{"0", "-10", "5", "null", "-3", "null", "9"}), // 一个可能的答案
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", treenode.Marshal(got), treenode.Marshal(tt.want))
			}
		})
	}
}
