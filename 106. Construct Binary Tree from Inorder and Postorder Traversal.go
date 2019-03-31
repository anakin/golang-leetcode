package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
根据中序和后续遍历结果，重建二叉树
后续遍历的最后一个元素，是二叉树的root节点
到中序数组中，root节点的位置，则左边用来递归创建左子树，右边用来递归创建右子树
*/

type Index struct {
	index int
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	//后续遍历的最后一个元素，是二叉树的root节点
	rootIndex := &Index{len(postorder) - 1}
	return buildTreeR(inorder, postorder, 0, len(postorder)-1, rootIndex)
}

func buildTreeR(inorder []int, postorder []int, start int, end int, rootIndex *Index) *TreeNode {
	if start > end {
		return nil
	}

	rootValue := postorder[rootIndex.index]
	root := &TreeNode{rootValue, nil, nil}
	rootIndex.index--

	if start == end {
		return root
	}

	//找到中序数组中，root节点的位置，则左边用来递归创建左子树，右边用来递归创建右子树
	index := 0
	for i := start; i <= end; i++ {
		if inorder[i] == rootValue {
			index = i
			break
		}
	}

	root.Right = buildTreeR(inorder, postorder, index+1, end, rootIndex)
	root.Left = buildTreeR(inorder, postorder, start, index-1, rootIndex)

	return root
}
