/*
4. Бинарное дерево в большую сумму дерева.
Дано root звено бинарного поискового дерева (BST), преобразуйте его в большее дерево так, чтобы каждый ключ исходного BST был изменен на исходный ключ плюс сумма всех ключей, больших исходного ключа в BST.

В качестве напоминания, бинарное поисковое дерево - это дерево, которое удовлетворяет следующим ограничениям:

Левое поддерево узла содержит только узлы с ключами меньше ключа узла.
Правое поддерево узла содержит только узлы с ключами больше ключа узла.
Оба поддерева должны быть также бинарными поисковыми деревьями.
Пример 1:

Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]

Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

Пример 2:

Input: root = [0,null,1]

Output: [1,null,1]

Constraints:

Диапазон значений узлов находится в диапазоне [1, 100]
0 <= Node.val <= 100
Все значения в дереве уникальны.
https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
*/

package main

func bstToGst(root *TreeNode) *TreeNode {
	total := 0
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {

		if node == nil {
			return total
		}

		node.Val += dfs(node.Right)
		total = node.Val
		dfs(node.Left)

		return total
	}

	_ = dfs(root)
	return root

}
