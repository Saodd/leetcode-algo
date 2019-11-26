package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var printCache [][]string

func PrintTreeNodes(root *TreeNode) {
	printCache = [][]string{}
	printTreeNodes(root, 0)
	for i := range printCache {
		fmt.Println(printCache[i])
	}
}

func printTreeNodes(node *TreeNode, depth int) {
	if len(printCache) <= depth {
		printCache = append(printCache, []string{})
	}
	if node == nil {
		printCache[depth] = append(printCache[depth], "null")
		return
	}
	printCache[depth] = append(printCache[depth], fmt.Sprint(node.Val))
	printTreeNodes(node.Left, depth+1)
	printTreeNodes(node.Right, depth+1)
}
