package q0021

/*
2021-03-23
执行用时：4 ms, 在所有 Go 提交中击败了67.60%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了97.82%的用户
*/
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	var cur = head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
