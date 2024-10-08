/*
6. Максимальное бинарное дерево
Вам дан массив целых чисел nums без дубликатов. Максимальное бинарное дерево может быть построено рекурсивно из nums с помощью следующего алгоритма:

Создайте корневой узел, значение которого является максимальным значением в nums.

Рекурсивно постройте левое поддерево на подмассиве префикса левее максимального значения.

Рекурсивно постройте правое поддерево на подмассиве суффикса правее максимального значения.

Верните максимальное бинарное дерево, построенное из nums.

Пример 1:

Input: nums = [3,2,1,6,0,5]

Output: [6,3,5,null,2,0,null,null,1]

Объяснение: Рекурсивные вызовы выглядят следующим образом:

- Наибольшее значение в [3,2,1,6,0,5] является 6. Левый префикс равен [3,2,1] и правый суффикс равен [0,5].

- Наибольшее значение в [3,2,1] является 3. Левый префикс равен [] и правый суффикс равен [2,1].

- Пустой массив, поэтому нет детей.

- Наибольшее значение в [2,1] является 2. Левый префикс равен [] и правый суффикс равен [1].

- Пустой массив, поэтому нет детей.

- Только один элемент, поэтому ребенок - узел со значением 1.

- Наибольшее значение в [0,5] является 5. Левый префикс равен [0] и правый суффикс равен [].

- Только один элемент, поэтому ребенок - узел со значением 0.

- Пустой массив, поэтому нет детей.

Пример 2:

Input: nums = [3,2,1]

Output: [3,null,2,null,1]

Constraints:

1 <= nums.length <= 1000
0 <= nums[i] <= 1000
All integers in nums are unique.
https://leetcode.com/problems/maximum-binary-tree/
*/

package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max, idx := maxNum(nums)
	return &TreeNode{max, constructMaximumBinaryTree(nums[:idx]), constructMaximumBinaryTree(nums[idx+1:])}
}

func maxNum(nums []int) (max, idx int) {
	for i, val := range nums {
		if i == 0 || val > max {
			max = val
			idx = i
		}
	}
	return
}
