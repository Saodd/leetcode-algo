package q0118

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "官方用例",
			args: args{5},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.numRows); !isEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if !reflect.DeepEqual(aa, b[i]) {
			return false
		}
	}
	return true
}
