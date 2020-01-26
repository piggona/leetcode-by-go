package flattenBinaryTreeToLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	if root.Left != nil {
		temp := root.Right
		root.Right = root.Left
		rtemp := root.Right
		for rtemp.Right != nil {
			rtemp = rtemp.Right
		}
		rtemp.Right = temp
	}
}

func flattenIter(root *TreeNode) {
	cur := root
	for cur.Left != nil || cur.Right != nil {
		if cur.Left != nil {
			temp := cur.Left
			right := cur.Right
			cur.Right = temp
			for temp.Right != nil {
				temp = temp.Right
			}
			temp.Right = right
		}
		cur = cur.Right
	}
}
