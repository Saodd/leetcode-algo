package p005x

import (
	"reflect"
	"testing"
)

func TestQuickSortInterval(t *testing.T) {
	type args struct {
		li [][]int
	}
	tests := []struct {
		name  string
		args  args
		wants [][]int
	}{
		{
			args:  args{[][]int{{1, 2}, {3, 4}, {2, 3}}},
			wants: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSortInterval(tt.args.li)
			if !reflect.DeepEqual(tt.args.li, tt.wants) {
				t.Fatal(tt.args.li, tt.wants)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			args: args{[][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
