package linked

/**
单链部分倒置，倒置第m到n个，
例如：m=2，n=5  输入 1-> 2->3->4->5 ->6->7
输出 1-> 5->4->3->2 ->6->7

92
https://leetcode.com/problems/reverse-linked-list-ii/
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	stub := new(ListNode)
	stub.Next = head
	var pmNode, mNode, nNode *ListNode = stub, nil, nil
	i := 0
	curr := stub
	for ; i<=m-1; i++ {
		if i == m-1 {
			pmNode = curr
		}
		curr = curr.Next
	}
	mNode = curr
	var prev, next *ListNode
	for ; i<=n; i++ {
		if i == n {
			nNode = curr
			next = curr.Next
			curr.Next = prev
			mNode.Next = next
			pmNode.Next = nNode
		}
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return stub.Next
}
