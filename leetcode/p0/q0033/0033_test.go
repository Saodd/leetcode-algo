package q0033

import (
	"testing"
)

type Solution func(nums []int, target int) int

func Test_indexOfMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}},
			want: 4,
		},
		{
			args: args{[]int{4, 5, 6, 0, 1, 2}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOfMin(tt.args.nums); got != tt.want {
				t.Errorf("indexOfMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{1, 2, 3, 4, 5}, 3},
			want: 2,
		},
		{
			args: args{[]int{1, 2, 3, 4}, 3},
			want: 2,
		},
		{
			args: args{[]int{1, 2, 3}, 3},
			want: 2,
		},
		{
			args: args{[]int{1, 2, 3, 4, 5}, 6},
			want: -1,
		},
		{
			args: args{[]int{1, 2, 3, 11, 12}, 6},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

type args struct {
	nums   []int
	target int
}
type Case struct {
	name string
	args args
	want int
}

func buildTests() []Case {
	return []Case{
		{
			args: args{[]int{1, 3}, 0},
			want: -1,
		},
		{
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			want: 4,
		},
		{
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			want: -1,
		},
		{
			name: "错误1",
			args: args{[]int{1}, 1},
			want: 0,
		},
		{
			name: "错误2",
			args: args{[]int{1, 3}, 1},
			want: 0,
		},
		{
			name: "错误3",
			args: args{[]int{1, 3}, 3},
			want: 1,
		},
		{
			name: "错误4",
			args: args{[]int{1, 3, 5}, 1},
			want: 0,
		},
	}
}

func do(t *testing.T, f Solution) {
	tests := buildTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search(t *testing.T) {
	do(t, search)
}
func Test_search2(t *testing.T) {
	do(t, search2)
}
