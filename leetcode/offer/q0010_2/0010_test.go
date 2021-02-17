package q0010_2

import "testing"

func Test_numWays(t *testing.T) {
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
			want: 2,
		},
		{
			name: "官方用例2",
			args: args{7},
			want: 21,
		},
		{
			name: "官方用例1",
			args: args{0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
