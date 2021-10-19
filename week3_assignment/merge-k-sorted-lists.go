package week3_assignment

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)

	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	middlePoint := length / 2
	leftPart := mergeKLists(lists[:middlePoint])
	rightPart := mergeKLists(lists[middlePoint:length])

	return mergeTwoLists(leftPart, rightPart)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	head := result

	for {
		if l1 == nil || l2 == nil {
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
