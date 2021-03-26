package treenode

import (
	"reflect"
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "用例1:全左",
			args: args{&TreeNode{Left: &TreeNode{Left: &TreeNode{}}}},
			want: []string{"0", "0", "null", "0"},
		},
		{
			name: "用例2:右左",
			args: args{&TreeNode{Right: &TreeNode{Left: &TreeNode{}}}},
			want: []string{"0", "null", "0", "null", "null", "0"},
		},
		{
			name: "用例3:空",
			args: args{nil},
			want: []string{},
		},
		{
			name: "用例4:单个",
			args: args{&TreeNode{}},
			want: []string{"0"},
		},
		{
			name: "用例5:全右",
			args: args{&TreeNode{Right: &TreeNode{Right: &TreeNode{}}}},
			want: []string{"0", "null", "0", "null", "null", "null", "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTree(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name     string
		args     args
		wantRoot *TreeNode
	}{
		{
			name:     "用例1",
			args:     args{[]string{"0", "0", "null", "0", "null", "null", "null"}},
			wantRoot: &TreeNode{Left: &TreeNode{Left: &TreeNode{}}},
		},
		{
			name:     "用例1(短)",
			args:     args{[]string{"0", "0", "null", "0"}},
			wantRoot: &TreeNode{Left: &TreeNode{Left: &TreeNode{}}},
		},
		{
			name:     "用例2",
			args:     args{[]string{}},
			wantRoot: nil,
		},
		{
			name:     "用例3",
			args:     args{[]string{"2", "1", "3"}},
			wantRoot: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		},
		{
			name: "用例4",
			args: args{[]string{"1", "2", "3", "4", "5", "6"}},
			wantRoot: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				}},
		},
		{
			name:     "互相测试1",
			args:     args{Format(NewTree([]string{"0", "0", "null", "0", "null", "null", "null"}))},
			wantRoot: &TreeNode{Left: &TreeNode{Left: &TreeNode{}}},
		},
		{
			name:     "互相测试2",
			args:     args{[]string{"0", "0", "null", "0", "null", "null", "null"}},
			wantRoot: NewTree(Format(&TreeNode{Left: &TreeNode{Left: &TreeNode{}}})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := NewTree(tt.args.words); !reflect.DeepEqual(gotRoot, tt.wantRoot) {
				t.Errorf("NewTree() = %v, want %v", gotRoot, tt.wantRoot)
			}
		})
	}
}
