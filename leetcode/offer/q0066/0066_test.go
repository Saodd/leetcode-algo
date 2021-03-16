package q0066

import (
	"reflect"
	"testing"
)

func Test_constructArr(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "官方用例",
			args: args{
				a: []int{1, 2, 3, 4, 5},
			},
			want: []int{120, 60, 40, 30, 24},
		},
		{
			name: "小数组",
			args: args{
				a: []int{3, 4},
			},
			want: []int{4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructArr(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
