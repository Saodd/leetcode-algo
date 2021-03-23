package q0021

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

type args struct {
	l1 *ListNode
	l2 *ListNode
}
type Case struct {
	name string
	args args
	want *ListNode
}

func buildTests() []Case {
	return []Case{
		{
			args: args{common.CreateListChain([]int{1, 2, 3}), common.CreateListChain([]int{4, 5, 6})},
			want: common.CreateListChain([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			args: args{common.CreateListChain([]int{1, 5, 6}), common.CreateListChain([]int{2, 3, 4})},
			want: common.CreateListChain([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			args: args{common.CreateListChain([]int{1, 2, 3}), common.CreateListChain([]int{})},
			want: common.CreateListChain([]int{1, 2, 3}),
		},
		{
			args: args{common.CreateListChain([]int{}), common.CreateListChain([]int{})},
			want: common.CreateListChain([]int{}),
		},
	}
}

func do(t *testing.T, f Solution) {
	for _, tt := range buildTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	do(t, mergeTwoLists)
}

func Test_mergeTwoLists2(t *testing.T) {
	do(t, mergeTwoLists2)
}
