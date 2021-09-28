//合并两个有序链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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


//加一
func plusOne(digits []int) []int {
    len := len(digits)
        var jinwei bool 

    for i := len - 1; i >= 0; i-- {
        if jinwei == true || i == len - 1 {
            if digits[i] + 1 > 9{
                jinwei = true
                digits[i] = 0

                if i == 0 {
                    temp := []int{1}
                    temp = append(temp, digits...)
                    digits = temp
                    break
                }
            } else {
                digits[i] += 1
                break
            }
        } 
    }

    return digits
}
