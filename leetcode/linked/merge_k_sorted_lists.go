package linked
/*
 * Descriptions:
 * https://leetcode.com/problems/merge-k-sorted-lists/
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	return partition(lists, 0, len(lists)-1)
}

func partition(lists []*ListNode, s int, e int)  *ListNode {
	if s == e {
		return lists[s]
	}
	if (s < e) {
		mid := (s+e)/2
		l1 := partition(lists, s, mid)
		l2 := partition(lists, mid+1, e)
		return mergeTwoLists(l1, l2)
	} else {
		return nil
	}
}