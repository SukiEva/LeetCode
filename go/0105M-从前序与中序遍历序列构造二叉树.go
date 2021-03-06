/*
 * @Author: SukiEva
 * @Date: 2021-11-10 19:12:09
 * @Description: https://github.com/SukiEva
 * わたし、気になります！
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return root
	}

	var rootIndex int
	for i, v := range inorder {
		if v == preorder[0] {
			rootIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}
