package q0148

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
	"reflect"
	"testing"
)

type Solution func(head *ListNode) *ListNode
type args struct {
	head *ListNode
}
type Case struct {
	name string
	args args
	want *ListNode
}

func BuildTests() []Case {
	return []Case{
		{
			name: "示例1",
			args: args{listnode.NewList([]int{4, 2, 1, 3})},
			want: listnode.NewList([]int{1, 2, 3, 4}),
		},
		{
			name: "示例2",
			args: args{listnode.NewList([]int{-1, 5, 3, 4, 0})},
			want: listnode.NewList([]int{-1, 0, 3, 4, 5}),
		},
		{
			name: "顺序",
			args: args{listnode.NewList([]int{1, 2, 3, 4})},
			want: listnode.NewList([]int{1, 2, 3, 4}),
		},
		{
			name: "倒序",
			args: args{listnode.NewList([]int{4, 3, 2, 1})},
			want: listnode.NewList([]int{1, 2, 3, 4}),
		},
		{
			name: "大一点",
			args: args{listnode.NewList([]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9})},
			want: listnode.NewList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
	}
}

func Do(t *testing.T, f Solution) {
	tests := BuildTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", listnode.Format(got), listnode.Format(tt.want))
			}
		})
	}
}

func Test_sortList2(t *testing.T) {
	Do(t, sortList2)
}
func Test_sortList1(t *testing.T) {
	Do(t, sortList1)
}
