package q0031

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}
type Case struct {
	name string
	args args
	want []int
}

func buildTests() []Case {
	return []Case{
		{
			args: args{[]int{1, 2, 3}},
			want: []int{1, 3, 2},
		},
		{
			args: args{[]int{1, 3, 2}},
			want: []int{2, 1, 3},
		},
		{
			args: args{[]int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			args: args{[]int{1, 1, 5}},
			want: []int{1, 5, 1},
		},
		{
			args: args{[]int{1, 2, 4, 5, 3}},
			want: []int{1, 2, 5, 3, 4},
		},
		{
			args: args{[]int{1, 2, 7, 4, 3, 6, 5}},
			want: []int{1, 2, 7, 4, 5, 3, 6},
		},
		{
			args: args{[]int{1, 2, 7, 4, 6, 5, 3}},
			want: []int{1, 2, 7, 5, 3, 4, 6},
		},
		{
			args: args{[]int{}},
			want: []int{},
		},
		{
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "错误",
			args: args{[]int{5, 4, 7, 5, 3, 2}},
			want: []int{5, 5, 2, 3, 4, 7},
		},
	}
}

func do(t *testing.T, f func([]int)) {
	for _, tt := range buildTests() {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("NextPermutation() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_nextPermutation1(t *testing.T) {
	do(t, nextPermutation1)
}

func Test_nextPermutation2(t *testing.T) {
	do(t, nextPermutation2)
}
func Test_nextPermutation3(t *testing.T) {
	do(t, nextPermutation3)
}
