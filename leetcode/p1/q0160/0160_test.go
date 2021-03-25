package q0160

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
	"reflect"
	"testing"
)

type Solution func(headA, headB *ListNode) *ListNode
type args struct {
	headA *ListNode
	headB *ListNode
}
type Case struct {
	name string
	args args
	want *ListNode
}

func BuildCase() []Case {
	l1 := listnode.Unmarshal([]int{1})
	l2 := listnode.Unmarshal([]int{1, 2})
	l5 := listnode.Unmarshal([]int{1, 2, 3, 4, 5})
	l6 := listnode.Unmarshal([]int{1, 2, 3, 4, 5, 6})

	h1a := listnode.Concat(listnode.Unmarshal([]int{100}), l1)
	h1b := listnode.Concat(listnode.Unmarshal([]int{200, 201}), l1)
	h2a := listnode.Concat(listnode.Unmarshal([]int{100}), l2)
	h2b := listnode.Concat(listnode.Unmarshal([]int{200, 201}), l2)
	h3a := listnode.Concat(listnode.Unmarshal([]int{100, 101, 102, 103}), l5)
	h3b := listnode.Concat(listnode.Unmarshal([]int{200, 201, 202}), l5)
	h4a := listnode.Concat(listnode.Unmarshal([]int{100, 101, 102, 103}), l6)
	h4b := listnode.Concat(listnode.Unmarshal([]int{200, 201, 202}), l6)
	h5a := listnode.Concat(listnode.Unmarshal([]int{100, 101, 102, 103, 104}), l6)
	h5b := listnode.Concat(listnode.Unmarshal([]int{200, 201, 202, 203, 204, 205}), l6)
	return []Case{
		{
			name: "自测1",
			args: args{h1a, h1b},
			want: l1,
		},
		{
			name: "自测2",
			args: args{h2a, h2b},
			want: l2,
		},
		{
			name: "自测3",
			args: args{h3a, h3b},
			want: l5,
		},
		{
			name: "自测4",
			args: args{h4a, h4b},
			want: l6,
		},
		{
			name: "自测5",
			args: args{h5a, h5b},
			want: l6,
		},
	}
}
func Do(t *testing.T, f Solution) {
	tests := BuildCase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntersectionNode2(t *testing.T) {
	Do(t, getIntersectionNode2)
}

func Test_getIntersectionNode1(t *testing.T) {
	Do(t, getIntersectionNode1)
}
