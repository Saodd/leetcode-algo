package strings

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{"Hello world"},
			want: 5,
		},
		{
			name: "空字符串",
			args: args{""},
			want: 0,
		},
		{
			name: "尾部空格",
			args: args{"Hello world "},
			want: 5,
		},
		{
			name: "开头空格",
			args: args{" world"},
			want: 5,
		},
		{
			name: "没有空格",
			args: args{"world"},
			want: 5,
		},
		{
			name: "没有字符",
			args: args{"    "},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
