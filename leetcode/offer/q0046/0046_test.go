package q0046

import "testing"

func Test_translateNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1位数",
			args: args{num: 1},
			want: 1,
		},
		{
			name: "2位数",
			args: args{num: 10},
			want: 2,
		},
		{
			name: "2位数",
			args: args{num: 26},
			want: 1,
		},
		{
			name: "3位数",
			args: args{num: 100},
			want: 2,
		},
		{
			name: "3位数",
			args: args{num: 101},
			want: 2,
		},
		{
			name: "3位数",
			args: args{num: 111},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
