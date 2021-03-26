package q0095

import (
	"github.com/Saodd/leetcode-algo/helper/treenode"
	"strings"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "官方用例",
			args: args{3},
			want: [][]string{
				{"1", "null", "2", "null", "null", "null", "3"},
				{"1", "null", "3", "null", "null", "2"},
				{"2", "1", "3"},
				{"3", "2", "null", "1"},
				{"3", "1", "null", "null", "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTrees(tt.args.n)
			compare := make(map[string]int)
			for _, tree := range got {
				compare[strings.Join(treenode.Format(tree), ",")]++
			}
			for _, words := range tt.want {
				compare[strings.Join(words, ",")]--
			}
			for k, v := range compare {
				if v > 0 {
					t.Errorf("%s  (got %d)", k, v)
				} else if v < 0 {
					t.Errorf("%s  (want %d)", k, v)
				}
			}
		})
	}
}
