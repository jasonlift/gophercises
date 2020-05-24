package linked

/**
206
https://leetcode.com/problems/reverse-linked-list/
 */

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev, next *ListNode = nil, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// in my version, every node is new & copy
func reverseListSelf(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := &ListNode {
		Val: head.Val,
		Next: nil,
	}
	tail := head.Next
	for tail != nil {
		tmp := &ListNode {
			Val: tail.Val,
			Next: curr,
		}
		tail = tail.Next
		curr = tmp
	}
	return curr
}
