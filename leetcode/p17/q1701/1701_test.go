package q1701

import (
	"math"
	"testing"
)

func Test_averageWaitingTime(t *testing.T) {
	type args struct {
		customers [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "官方用例1",
			args: args{[][]int{{1, 2}, {2, 5}, {4, 3}}},
			want: 5,
		},
		{
			name: "官方用例2",
			args: args{[][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}},
			want: 3.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageWaitingTime(tt.args.customers); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("averageWaitingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
