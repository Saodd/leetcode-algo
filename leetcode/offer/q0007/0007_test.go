package q0007

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "官方用例1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.preorder, tt.args.inorder)
			if !isEqualTree(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func printTree(node *TreeNode) (res []int) {
	if node == nil {
		return nil
	}
	res = append(res, node.Val)
	res = append(res, printTree(node.Left)...)
	res = append(res, printTree(node.Right)...)
	return res
}

func isEqualTree(a, b *TreeNode) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	if !isEqualTree(a.Left, b.Left) || !isEqualTree(a.Right, b.Right) {
		return false
	}
	return true
}
