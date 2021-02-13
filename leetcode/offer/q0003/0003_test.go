package q0003

import "testing"

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "官方用例",
			args: args{[]int{2, 3, 1, 0, 2, 5, 3}},
			want: []int{2, 3},
		},
		{
			name: "暴力排序",
			args: args{[]int{2, 3, 1, 2}},
			want: []int{2},
		},
		{
			name: "错误1",
			args: args{[]int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
			want: []int{11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); !isIn(got, tt.want) {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isIn(x int, nums []int) bool {
	for _, v := range nums {
		if v == x {
			return true
		}
	}
	return false
}
