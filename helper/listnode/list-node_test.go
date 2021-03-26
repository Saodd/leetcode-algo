package listnode

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "用例1",
			args: args{[]int{1, 2, 3}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "用例2",
			args: args{nil},
			want: nil,
		},
		{
			name: "用例3",
			args: args{[]int{1}},
			want: &ListNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "用例1",
			args: args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
			want: []int{1, 2, 3},
		},
		{
			name: "用例2",
			args: args{nil},
			want: nil,
		},
		{
			name: "用例3",
			args: args{&ListNode{Val: 1}},
			want: []int{1},
		},
		{
			name: "互相调用",
			args: args{NewList([]int{1, 2, 3})},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	type args struct {
		front *ListNode
		back  *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "用例1",
			args: args{NewList([]int{1, 2, 3}), NewList([]int{4, 5, 6})},
			want: NewList([]int{1, 2, 3, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.front, tt.args.back); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
