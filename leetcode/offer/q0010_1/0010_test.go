package q0010_1

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "官方用例1",
			args: args{2},
			want: 1,
		},
		{
			name: "官方用例2",
			args: args{5},
			want: 5,
		},
		{
			name: "错误1，就很奇怪为什么要取余数？",
			args: args{45},
			want: 134903163,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
