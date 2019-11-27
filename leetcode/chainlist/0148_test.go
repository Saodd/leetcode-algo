package chainlist

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
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
			args: args{common.CreateListChain([]int{4, 2, 1, 3})},
			want: common.CreateListChain([]int{1, 2, 3, 4}),
		},
		{
			name: "示例2",
			args: args{common.CreateListChain([]int{-1, 5, 3, 4, 0})},
			want: common.CreateListChain([]int{-1, 0, 3, 4, 5}),
		},
		{
			name: "顺序",
			args: args{common.CreateListChain([]int{1, 2, 3, 4})},
			want: common.CreateListChain([]int{1, 2, 3, 4}),
		},
		{
			name: "倒序",
			args: args{common.CreateListChain([]int{4, 3, 2, 1})},
			want: common.CreateListChain([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				common.PrintListChain(tt.args.head)
				common.PrintListChain(tt.want)
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
