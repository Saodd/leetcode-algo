package chainlist

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
			args: args{common.CreateListChain([]int{1, 2, 3, 3, 4, 4, 5})},
			want: common.CreateListChain([]int{1, 2, 5}),
		},
		{
			name: "示例2",
			args: args{common.CreateListChain([]int{1, 1, 1, 2, 3})},
			want: common.CreateListChain([]int{2, 3}),
		},
		{
			name: "空链表",
			args: args{common.CreateListChain([]int{})},
			want: common.CreateListChain([]int{}),
		},
		{
			name: "易错点：尾部重复",
			args: args{common.CreateListChain([]int{1, 2, 2})},
			want: common.CreateListChain([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
				common.PrintListChain(got)
				common.PrintListChain(tt.want)
			}
		})
	}
}
