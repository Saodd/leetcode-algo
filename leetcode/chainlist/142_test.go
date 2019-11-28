package chainlist

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	chain1 := common.CreateListChainCycle([]int{3, 2, 0, -4}, 1)
	chain2 := common.CreateListChainCycle([]int{1, 2}, 0)
	chain3 := common.CreateListChainCycle([]int{1}, -1)

	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "示例1",
			args: args{head: chain1},
			want: chain1.Next,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
