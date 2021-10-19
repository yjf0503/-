package week3_assignment

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder)
}

func build(inorder, postorder []int) *TreeNode {
	if len(postorder) <= 0 {
		return nil
	}

	var node = TreeNode{
		Val: postorder[len(postorder)-1],
	}

	var midIndex int
	for k, v := range inorder {
		if v == node.Val {
			midIndex = k
		}
	}

	node.Left = build(inorder[:midIndex], postorder[:midIndex])
	node.Right = build(inorder[midIndex+1:], postorder[midIndex:len(postorder)-1])

	return &node
}
