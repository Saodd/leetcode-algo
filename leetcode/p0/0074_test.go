package p0

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "错误一",
			args: args{
				matrix: [][]int{{1}, {3}},
				target: 3,
			},
			want: true,
		},
		{
			name: "示例一",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 3,
			},
			want: true,
		},
		{
			name: "示例二",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 13,
			},
			want: false,
		},
		{
			name: "自定义：空数组",
			args: args{
				matrix: [][]int{{}},
				target: 3,
			},
			want: false,
		},
		{
			name: "自定义：超过边缘",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 51,
			},
			want: false,
		},
		{
			name: "自定义：命中行首",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 10,
			},
			want: true,
		},
		{
			name: "自定义：偶数行矩阵",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}, {51, 52, 53, 54}},
				target: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
			if got := searchMatrix2(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
