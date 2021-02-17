package q0011

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "官方用例1",
			args: args{[]int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "官方用例2",
			args: args{[]int{2, 2, 2, 0, 1}},
			want: 0,
		},
		{
			name: "预测易错点-1",
			args: args{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 2, 2, 2, 2, 2}},
			want: 0,
		},
		{
			name: "预测易错点-2",
			args: args{[]int{2, 2, 2, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
