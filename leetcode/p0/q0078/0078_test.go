package q0078

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "官方用例1",
			args: args{[]int{1, 2}},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "官方用例2",
			args: args{[]int{0}},
			want: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
