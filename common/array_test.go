package common

import "testing"

func TestIntArrayBinaryFind(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "长度1",
			args: args{
				list:   []int{1},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntArrayBinaryFind(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("IntArrayBinaryFind() = %v, want %v", got, tt.want)
			}
		})
	}
}
