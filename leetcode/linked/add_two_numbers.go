package linked

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := 0
	residue := 0
	var prev, curr *ListNode = &ListNode{
		Val: 0,
		Next: nil,
	}, nil
	head := prev

	for l1 != nil {
		if l2 != nil {
			tmp = l1.Val + l2.Val + residue
			curr = &ListNode {
				Val: tmp % 10,
				Next: nil,
			}
			prev.Next = curr
			prev = curr
			residue = tmp / 10

			l1, l2 = l1.Next, l2.Next
		} else {
			tmp = l1.Val+residue
			curr = &ListNode {
				Val: tmp % 10,
				Next: nil,
			}
			prev.Next = curr
			prev = curr
			residue = tmp / 10

			l1 = l1.Next
		}
	}
	for l2 != nil {
		tmp = l2.Val+residue
		curr = &ListNode {
			Val: tmp % 10,
			Next: nil,
		}
		prev.Next = curr
		prev = curr
		residue = tmp / 10

		l2 = l2.Next
	}
	if residue != 0 {
		prev.Next = &ListNode {
			Val: residue,
			Next: nil,
		}
	}
	return head.Next
}