package p3

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
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
			args: args{common.CreateListChain([]int{1, 2, 3, 4, 5})},
			want: common.CreateListChain([]int{1, 3, 5, 2, 4}),
		},
		{
			name: "示例2",
			args: args{common.CreateListChain([]int{2, 1, 3, 5, 6, 4, 7})},
			want: common.CreateListChain([]int{2, 3, 6, 7, 1, 5, 4}),
		},
		{
			name: "偶数长",
			args: args{common.CreateListChain([]int{1, 2, 3, 4, 5, 6})},
			want: common.CreateListChain([]int{1, 3, 5, 2, 4, 6}),
		},
		{
			name: "空",
			args: args{common.CreateListChain([]int{})},
			want: common.CreateListChain([]int{}),
		},
		{
			name: "短",
			args: args{common.CreateListChain([]int{1})},
			want: common.CreateListChain([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
