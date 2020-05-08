package p0

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例一",
			args: args{nums: []int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "自定义：无需排序",
			args: args{nums: []int{0, 0, 1, 1, 2, 2, 2}},
			want: []int{0, 0, 1, 1, 2, 2, 2},
		},
		{
			name: "自定义：全0",
			args: args{nums: []int{0, 0, 0}},
			want: []int{0, 0, 0},
		},
		{
			name: "自定义：全1",
			args: args{nums: []int{1, 1, 1}},
			want: []int{1, 1, 1},
		},
		{
			name: "自定义：全2",
			args: args{nums: []int{2, 2, 2}},
			want: []int{2, 2, 2},
		},
		{
			name: "错误1",
			args: args{nums: []int{1, 2, 0}},
			want: []int{0, 1, 2},
		},
		{
			name: "错误2",
			args: args{nums: []int{2, 0}},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("got = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
