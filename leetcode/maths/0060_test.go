package maths

import (
	"reflect"
	"testing"
)

func Test_getPermutation(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "示例1",
			args: args{3, 3},
			want: "213",
		},
		{
			name: "示例2",
			args: args{4, 9},
			want: "2314",
		},
		{
			name: "小数1",
			args: args{1, 1},
			want: "1",
		},
		{
			name: "小数2",
			args: args{2, 2},
			want: "21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upInsertBytes(t *testing.T) {
	type args struct {
		tar int
		ori int
		bs  []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{0, 0, []byte{'1', '2', '3'}},
			want: []byte{'1', '2', '3'},
		},
		{
			args: args{0, 0, []byte{'1', '2', '3'}},
			want: []byte{'1', '2', '3'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if upInsertBytes(tt.args.tar, tt.args.ori, tt.args.bs); !reflect.DeepEqual(tt.args.bs, tt.want) {
				t.Errorf("getPermutation() = %v, want %v", tt.args.bs, tt.want)
			}
		})
	}
}
