package p0

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *common.ListNode
		l2 *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
