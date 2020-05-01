package p0

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例一",
			args: args{matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			want: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name: "示例二",
			args: args{matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}},
			want: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			name: "错误一：在首行或首列有0的情况",
			args: args{matrix: [][]int{{1, 1, 1}, {0, 1, 2}}},
			want: [][]int{{0, 1, 1}, {0, 0, 0}},
		},
		{
			name: "错误二：typo",
			args: args{matrix: [][]int{{0, 1}}},
			want: [][]int{{0, 0}},
		},
		{
			name: "错误三：typo",
			args: args{matrix: [][]int{{3, 5, 5, 6, 9, 1, 4, 5, 0, 5}, {2, 7, 9, 5, 9, 5, 4, 9, 6, 8}, {6, 0, 7, 8, 1, 0, 1, 6, 8, 1}, {7, 2, 6, 5, 8, 5, 6, 5, 0, 6}, {2, 3, 3, 1, 0, 4, 6, 5, 3, 5}, {5, 9, 7, 3, 8, 8, 5, 1, 4, 3}, {2, 4, 7, 9, 9, 8, 4, 7, 3, 7}, {3, 5, 2, 8, 8, 2, 2, 4, 9, 8}}},
			want: [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {2, 0, 9, 5, 0, 0, 4, 9, 0, 8}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {5, 0, 7, 3, 0, 0, 5, 1, 0, 3}, {2, 0, 7, 9, 0, 0, 4, 7, 0, 7}, {3, 0, 2, 8, 0, 0, 2, 4, 0, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if setZeroes(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
