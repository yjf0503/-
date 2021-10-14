package week1_assignment

// ListNode
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	head := result

	for {
		if l1 == nil {
			break
		}

		if l2 == nil {
			break
		}

		if l1.Val < l2.Val {
			result.Next = l1
			l1 = l1.Next
		} else {
			result.Next = l2
			l2 = l2.Next
		}
		result = result.Next
	}

	if l1 == nil {
		result.Next = l2
	}

	if l2 == nil {
		result.Next = l1
	}

	return head.Next
}
