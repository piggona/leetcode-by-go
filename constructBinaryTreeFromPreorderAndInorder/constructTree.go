package constructBinaryTreeFromPreorderAndInorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	} else {
		target := preorder[0]
		var ltemp, rtemp []int
		var lpre, rpre []int
		for i := 0; i < len(inorder); i++ {
			if inorder[i] == target {
				ltemp = inorder[:i]
				rtemp = inorder[i+1:]
				lpre = preorder[1 : i+1]
				rpre = preorder[i+1:]
			}
		}
		return &TreeNode{
			Val:   preorder[0],
			Left:  buildTree(lpre, ltemp),
			Right: buildTree(rpre, rtemp),
		}
	}
}
