package q0142

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	chain1 := listnode.NewListWithCycle([]int{3, 2, 0, -4}, 1)
	chain2 := listnode.NewListWithCycle([]int{1, 2}, 0)
	chain3 := listnode.NewListWithCycle([]int{1}, -1)
	chain4 := listnode.NewListWithCycle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9)

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例1",
			args: args{head: chain1},
			want: listnode.Walk(chain1, 1),
		},
		{
			name: "示例2",
			args: args{head: chain2},
			want: chain2,
		},
		{
			name: "示例3",
			args: args{head: chain3},
			want: nil,
		},
		{
			name: "自测1",
			args: args{head: chain4},
			want: listnode.Walk(chain4, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
