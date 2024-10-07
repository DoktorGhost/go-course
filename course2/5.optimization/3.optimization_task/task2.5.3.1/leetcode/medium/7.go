/*
7. Балансировка бинарного поискового дерева
Дано корневое значение бинарного поискового дерева. Верните сбалансированное бинарное поисковое дерево с теми же значениями узлов. Если есть более одного ответа, верните любой из них.

Бинарное поисковое дерево считается сбалансированным, если глубина двух поддеревьев каждого узла никогда не отличается более чем на 1.

Пример 1:

Input: root = [1,null,2,null,3,null,4,null,null]

Output: [2,1,3,null,null,null,4]

Explanation: Это не единственный правильный ответ, [3,1,4,null,2] также правильный.

Пример 2:

Input: root = [2,1,3]

Output: [2,1,3]

Constraints:

Количество узлов в дереве находится в диапазоне [1, 104].
1 <= Node.val <= 105
https://leetcode.com/problems/balance-a-binary-search-tree/
*/
package main

func balanceBST(root *TreeNode) *TreeNode {
	nodes := []int{}
	inorderTraversal(root, &nodes)
	return sortedArrayToBST(nodes)
}

func inorderTraversal(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, nodes)
	*nodes = append(*nodes, root.Val)
	inorderTraversal(root.Right, nodes)
}

func sortedArrayToBST(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	mid := len(nodes) / 2
	root := &TreeNode{Val: nodes[mid]}
	root.Left = sortedArrayToBST(nodes[:mid])
	root.Right = sortedArrayToBST(nodes[mid+1:])
	return root
}
