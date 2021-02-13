package q0079

import "testing"

func Test_exist(t *testing.T) {
	theBoard := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "官方用例1",
			args: args{
				board: theBoard,
				word:  "ABCCED",
			},
			want: true,
		},
		{
			name: "官方用例2",
			args: args{
				board: theBoard,
				word:  "SEE",
			},
			want: true,
		},
		{
			name: "官方用例3",
			args: args{
				board: theBoard,
				word:  "ABCB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
