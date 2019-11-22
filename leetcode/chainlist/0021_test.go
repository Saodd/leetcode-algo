package chainlist

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
			args: args{common.CreateListInt([]int{1, 2, 3}), common.CreateListInt([]int{4, 5, 6})},
			want: common.CreateListInt([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			args: args{common.CreateListInt([]int{1, 5, 6}), common.CreateListInt([]int{2, 3, 4})},
			want: common.CreateListInt([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			args: args{common.CreateListInt([]int{1, 2, 3}), common.CreateListInt([]int{})},
			want: common.CreateListInt([]int{1, 2, 3}),
		},
		{
			args: args{common.CreateListInt([]int{}), common.CreateListInt([]int{})},
			want: common.CreateListInt([]int{}),
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
