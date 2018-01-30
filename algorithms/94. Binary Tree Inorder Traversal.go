package algorithms

//Given a binary tree, return the inorder traversal of its nodes' values.
//
//For example:
//Given binary tree [1,null,2,3],
//1
//\
//2
///
//3
//return [1,3,2].
//
//Note: Recursive solution is trivial, could you do it iteratively?

func inorderTraversal(root *TreeNode) (r []int) {
	inorderTraversalHelper(root, &r)
	return
}

func inorderTraversalHelper(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorderTraversalHelper(root.Left, r)
	}
	*r = append(*r, root.Val)
	if root.Right != nil {
		inorderTraversalHelper(root.Right, r)
	}
}
