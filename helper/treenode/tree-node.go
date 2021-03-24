package treenode

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Marshal(t *TreeNode) []string {
	var res []string
	layer := []*TreeNode{t}
	next := true
	for next {
		next = false
		nextLayer := make([]*TreeNode, len(layer)*2)
		for i, parent := range layer {
			if parent == nil {
				res = append(res, "null")
				nextLayer[i*2], nextLayer[i*2+1] = nil, nil
			} else {
				res = append(res, strconv.Itoa(parent.Val))
				nextLayer[i*2], nextLayer[i*2+1] = parent.Left, parent.Right
				if parent.Left != nil || parent.Right != nil {
					next = true
				}
			}
		}
		layer = nextLayer
	}
	{
		var cut int
		for cut = len(res); cut > 0 && res[cut-1] == "null"; cut-- {
		}
		if cut >= 0 {
			res = res[:cut]
		}
	}
	return res
}

func Unmarshal(words []string) (root *TreeNode) {
	if len(words) == 0 {
		return
	}
	{ // 为了取得root，对第一层特殊处理一下
		value, err := strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}
		root = &TreeNode{Val: value}
		words = words[1:]
	}
	layer := []*TreeNode{root}
	for width := 2; ; width *= 2 { // 从第二层开始循环
		if len(words) == 0 {
			return
		}
		var layerWords []string
		if len(words) < width {
			layerWords = words
		} else {
			layerWords = words[:width]
		}

		nextLayer := make([]*TreeNode, width)
		for i, word := range layerWords {
			if word != "null" {
				value, err := strconv.Atoi(word)
				if err != nil {
					panic(err)
				}
				nextLayer[i] = &TreeNode{Val: value}
			}
		}

		for i, parent := range layer {
			if parent != nil {
				parent.Left = nextLayer[i*2]
				parent.Right = nextLayer[i*2+1]
			}
		}

		layer = nextLayer
		if len(words) < width {
			words = words[:0]
		} else {
			words = words[:width]
		}
	}
}
