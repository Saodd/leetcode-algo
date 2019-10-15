package common

import (
	"math/rand"
	"testing"
)

func genIntList(num int) []int {
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
			args: args{genIntList(2000000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
