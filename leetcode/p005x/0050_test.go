package p005x

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "示例",
			args: args{2.00000, 10},
			want: 1024.0,
		},
		{
			name: "示例",
			args: args{2.10000, 3},
			want: 9.26100,
		},
		{
			name: "示例",
			args: args{2.00000, -2},
			want: 0.25,
		},
		{
			name: "负数x",
			args: args{-2.00000, -2},
			want: 0.25,
		},
		{
			name: "大n",
			args: args{2, 18},
			want: 262144,
		},
		{
			name: "负数大n",
			args: args{8.84372, -5},
			want: 2e-05,
		},
		{
			name: "N的整形边界",
			args: args{1.00000000001, -1 << 31},
			want: 0.97875410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n) - tt.want; !(got < 1e-5 && got > -1e-5) {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
