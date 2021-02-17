package q0013

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "官方用例1",
			args: args{2, 3, 1},
			want: 3,
		},
		{
			name: "官方用例2",
			args: args{3, 1, 0},
			want: 1,
		},
		{
			name: "错误-1",
			args: args{16, 8, 4},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
