package q0082

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
	"reflect"
	"testing"
)

type Solution func(*ListNode) *ListNode
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
			args: args{listnode.NewList([]int{1, 2, 3, 3, 4, 4, 5})},
			want: listnode.NewList([]int{1, 2, 5}),
		},
		{
			name: "示例2",
			args: args{listnode.NewList([]int{1, 1, 1, 2, 3})},
			want: listnode.NewList([]int{2, 3}),
		},
		{
			name: "空链表",
			args: args{listnode.NewList([]int{})},
			want: listnode.NewList([]int{}),
		},
		{
			name: "易错点：尾部重复",
			args: args{listnode.NewList([]int{1, 2, 2})},
			want: listnode.NewList([]int{1}),
		},
	}
}

func Do(t *testing.T, solution Solution) {
	tests := BuildTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listnode.Format(solution(tt.args.head))
			want := listnode.Format(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("deleteDuplicates() = %+v, want %v", got, want)
			}
		})
	}
}

func Test_deleteDuplicates3(t *testing.T) {
	Do(t, deleteDuplicates3)
}
func Test_deleteDuplicates1(t *testing.T) {
	Do(t, deleteDuplicates1)
}
func Test_deleteDuplicates2(t *testing.T) {
	Do(t, deleteDuplicates2)
}
