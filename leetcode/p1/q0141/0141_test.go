package q0141

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{listnode.NewListWithCycle([]int{3, 2, 0, -4}, 1)},
			want: true,
		},
		{
			name: "示例2",
			args: args{listnode.NewListWithCycle([]int{1, 2}, 0)},
			want: true,
		},
		{
			name: "示例3",
			args: args{listnode.NewListWithCycle([]int{1}, -1)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
