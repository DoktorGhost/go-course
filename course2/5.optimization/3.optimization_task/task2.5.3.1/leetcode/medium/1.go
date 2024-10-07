/*
1. Сумма самых глубоких листьев.
Дан root дерево бинарных данных, верните сумму значений его самых глубоких листьев.

Пример 1:

Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]

Output: 15

Пример 2:

Input: root = [,,8,,,,9,null,,,null,null,null,]

Output: 19

Constraints:

Количество узлов в дереве находится в диапазоне [1, 104].
1 <= Node.val <= 100
https://leetcode.com/problems/deepest-leaves-sum/
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	result := make(map[int]int)
	deep := 0
	process(deep, result, root)
	maxDeep := 0
	maxRes := 0

	for key, value := range result {
		if key >= maxDeep {
			maxDeep = key
			maxRes = value
		}
	}
	return maxRes
}

func process(deep int, maps map[int]int, root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		maps[deep] += root.Val
		return
	}
	if root.Left != nil && root.Right != nil {
		deep++
		process(deep, maps, root.Left)
		process(deep, maps, root.Right)
		return
	}
	if root.Right != nil {
		deep++
		process(deep, maps, root.Right)
		return
	}
	if root.Left != nil {
		deep++
		process(deep, maps, root.Left)
		return
	}
}
