package chainlist

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *common.ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "示例1",
			args: args{common.CreateListInt([]int{1, 2, 3, 4, 5}), 2, 4},
			want: common.CreateListInt([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "走到尾部",
			args: args{common.CreateListInt([]int{1, 2, 3, 4, 5}), 2, 5},
			want: common.CreateListInt([]int{1, 5, 4, 3, 2}),
		},
		{
			name: "从头开始",
			args: args{common.CreateListInt([]int{1, 2, 3, 4, 5}), 1, 4},
			want: common.CreateListInt([]int{4, 3, 2, 1, 5}),
		},
		{
			name: "全部翻转",
			args: args{common.CreateListInt([]int{1, 2, 3, 4, 5}), 1, 5},
			want: common.CreateListInt([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
				common.PrintListInt(got)
				common.PrintListInt(tt.want)
			}
		})
	}
}
