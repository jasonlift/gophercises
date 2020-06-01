package linked

type ListNode struct {
	Val int
	Next *ListNode
}

func Array2LinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	prev := &ListNode {
		Val: 0,
		Next: nil,
	}
	head := prev

	for _, item := range arr {
		tmp := &ListNode {
			Val: item,
			Next: nil,
		}
		prev.Next = tmp
		prev = prev.Next
	}
	return head.Next
}

func LinkedList2Array(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func verifyCorrect(l *ListNode, answer []int) bool {
	for _, item := range answer {
		if l == nil || l.Val != item {
			return false
		}
		l = l.Next
	}
	if l != nil {
		return false
	}
	return true
}

