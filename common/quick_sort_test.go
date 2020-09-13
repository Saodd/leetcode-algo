package common

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func genIntList(num int) []int {
	rand.Seed(time.Now().Unix())
	data := make([]int, num)
	for i := 0; i < num; i++ {
		data[i] = rand.Intn(100)
	}
	return data
}

func TestQuickSortInt(t *testing.T) {
	type args struct {
		li []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{genIntList(50)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want = make([]int, len(tt.args.li))
			copy(want, tt.args.li)
			sort.Ints(want)
			QuickSortInt(tt.args.li)
			if !reflect.DeepEqual(tt.args.li, want) {
				t.Error("失败")
			}
		})
	}
}
