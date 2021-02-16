package q0006

import (
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "官方用例1",
			args: args{&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			}},
			want: []int{2, 3, 1},
		},
		{
			name: "空",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
