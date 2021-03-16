package mySortInt

import (
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	CommonTest(BubbleSort, t)

}

func TestQuickSort(t *testing.T) {
	CommonTest(QuickSort, t)
}

func TestSelectSort(t *testing.T) {
	CommonTest(SelectSort, t)
}

func CommonTest(f SortFunc, t *testing.T) {
	type args struct {
		builder ArrayBuilder
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{builder: NewArrayBuilder(1)}},
		{args: args{builder: NewArrayBuilder(2)}},
		{args: args{builder: NewArrayBuilder(3)}},
		{args: args{builder: func() []int { return nil }}},
		{args: args{builder: func() []int { return []int{1} }}},
		{args: args{builder: func() []int { return []int{1, 2, 3, 4, 5} }}},
		{args: args{builder: func() []int { return []int{5, 4, 3, 2, 1} }}},
		{args: args{builder: func() []int { return []int{1, 2, 3, 2, 1} }}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := tt.args.builder()
			f(li)
			CheckArray(li, t)
		})
	}
}

func CheckArray(li []int, t *testing.T) {
	for i := 1; i < len(li); i++ {
		if li[i-1] > li[i] {
			t.Errorf("错误位置：%d, 前：%d, 后：%d, 数组：%v", i, li[i-1], li[i], li)
			return
		}
	}
}

type SortFunc func([]int)
type ArrayBuilder func() []int

func NewArrayBuilder(seed int64) ArrayBuilder {
	return func() []int {
		r := rand.New(rand.NewSource(seed))
		li := make([]int, 1000)
		for i := range li {
			li[i] = r.Intn(10000)
		}
		return li
	}
}
