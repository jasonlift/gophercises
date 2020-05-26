package linked

/**
142
https://leetcode.com/problems/linked-list-cycle-ii/
 */

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow2 := head
			for slow2 != fast {
				slow2 = slow2.Next
				fast = fast.Next
			}
			return slow2
		}
	}
	return nil
}