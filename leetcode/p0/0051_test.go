package p0

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "示例",
			args: args{4},
			want: [][]string{
				{"..Q.", "Q...", "...Q", ".Q.."},
				{".Q..", "...Q", "Q...", "..Q."},
			},
		},
		{
			name: "两个",
			args: args{2},
			want: [][]string{},
		},
		{
			name: "三个",
			args: args{3},
			want: [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
