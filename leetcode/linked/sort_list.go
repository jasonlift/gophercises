package linked

/**
148
https://leetcode.com/problems/sort-list/
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	slow, fast:= head, head
	var prev *ListNode = slow
	for fast != nil && fast.Next != nil{
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	l := sortList(head)
	r := sortList(slow)
	return merge(l, r)
}

func merge(l *ListNode, r *ListNode) *ListNode {
	tmpNode := &ListNode{
		Val: 0,
		Next: nil,
	}
	head := tmpNode
	for l != nil && r != nil {
		if l.Val < r.Val {
			tmpNode.Next = l
			l = l.Next
		} else {
			tmpNode.Next = r
			r = r.Next
		}
		tmpNode = tmpNode.Next
	}
	if l != nil {
		tmpNode.Next = l
	}
	if r != nil {
		tmpNode.Next = r
	}
	return head.Next
}