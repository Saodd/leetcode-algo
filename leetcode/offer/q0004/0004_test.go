package q0004

import "testing"

func Test_findNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	matrix2 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

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
			name: "官方用例1",
			args: args{
				matrix: matrix,
				target: 5,
			},
			want: true,
		},
		{
			name: "官方用例2",
			args: args{
				matrix: matrix,
				target: 20,
			},
			want: false,
		},
		{
			name: "边界条件-1",
			args: args{
				matrix: matrix,
				target: 0,
			},
			want: false,
		},
		{
			name: "边界条件-2",
			args: args{
				matrix: matrix,
				target: 1,
			},
			want: true,
		},
		{
			name: "边界条件-3",
			args: args{
				matrix: matrix,
				target: 15,
			},
			want: true,
		},
		{
			name: "边界条件-4",
			args: args{
				matrix: matrix,
				target: 100,
			},
			want: false,
		},
		{
			name: "边界条件-5",
			args: args{
				matrix: matrix,
				target: 30,
			},
			want: true,
		},
		{
			name: "边界条件-6",
			args: args{
				matrix: matrix,
				target: 31,
			},
			want: false,
		},
		{
			name: "错误-1",
			args: args{
				matrix: matrix2,
				target: 19,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
