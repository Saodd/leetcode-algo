package q0014_1

import "testing"

func Test_cuttingRope(t *testing.T) {
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
			args: args{10},
			want: 36,
		},
		{
			name: "自测-1",
			args: args{3},
			want: 2,
		},
		{
			name: "自测-2",
			args: args{4},
			want: 4,
		},
		{
			name: "自测-3",
			args: args{5},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
