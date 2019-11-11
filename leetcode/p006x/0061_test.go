package p006x

import (
	"github.com/Saodd/leetcode-algo/common"
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *common.ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "示例1",
			args: args{common.CreateListInt([]int{1, 2, 3, 4, 5}), 2},
			want: common.CreateListInt([]int{4, 5, 1, 2, 3}),
		},
		{
			name: "示例2",
			args: args{common.CreateListInt([]int{0, 1, 2}), 4},
			want: common.CreateListInt([]int{2, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
