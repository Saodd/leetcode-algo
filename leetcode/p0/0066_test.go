package p0

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{[]int{1, 2, 3}},
			want: []int{1, 2, 4},
		},
		{
			name: "示例2",
			args: args{[]int{4, 3, 2, 1}},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "进位1",
			args: args{[]int{1, 2, 9}},
			want: []int{1, 3, 0},
		},
		{
			name: "进位2",
			args: args{[]int{9, 9, 9}},
			want: []int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}