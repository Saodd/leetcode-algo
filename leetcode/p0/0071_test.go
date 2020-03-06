package p0

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "示例1",
			args: args{"/home/"},
			want: "/home",
		},
		{
			name: "示例2",
			args: args{"/../"},
			want: "/",
		},
		{
			name: "示例3",
			args: args{"/home//foo/"},
			want: "/home/foo",
		},
		{
			name: "示例4",
			args: args{"/a/./b/../../c/"},
			want: "/c",
		},
		{
			name: "示例5",
			args: args{"/a/../../b/../c//.//"},
			want: "/c",
		},
		{
			name: "示例6",
			args: args{"/a//b////c/d//././/.."},
			want: "/a/b/c",
		},
		{
			name: "错误1", // 这个很神奇，三个点就可以作为文件名或者目录名了
			args: args{"/..."},
			want: "/...",
		},
		{
			name: "错误2",
			args: args{"/.hidden"},
			want: "/.hidden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
