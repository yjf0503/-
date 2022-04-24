package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var step = 1
//	var queue = make([]*TreeNode, 0)
//	queue = append(queue, root)
//
//	for {
//		if len(queue) <= 0 {
//			break
//		}
//
//		size := len(queue)
//		for i := 0; i < size; i++ {
//			node := queue[0]
//			if node.Left == nil && node.Right == nil {
//				return step
//			}
//
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//			queue = queue[1:]
//		}
//
//		step++
//	}
//
//	return step
//}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getStep(root, 1)
}

func getStep(node *TreeNode, step int) int {
	if node == nil {
		return step
	}

	if node.Left == nil && node.Right == nil {
		return step
	}

	leftAns := getStep(node.Left, step+1)
	rightAns := getStep(node.Right, step+1)

	if node.Left == nil {
		return rightAns
	}

	if node.Right == nil {
		return leftAns
	}

	if leftAns < rightAns {
		return leftAns
	}

	return rightAns
}
