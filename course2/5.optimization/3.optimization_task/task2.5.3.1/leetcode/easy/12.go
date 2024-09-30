/*
12. Количество хороших пар
Дан массив целых чисел nums. Верните количество хороших пар.

Пара (i, j) называется хорошей, если nums [i] == nums [j] и i <j.

Пример 1:

Input: nums = [1,2,3,1,1,3]

Output: 4

Объяснение: Есть 4 хорошие пары (0,3), (0,4), (3,4), (2,5), с нулевым индексом.

Пример 2:

Input: nums = [1,1,1,1]

Output: 6

Объяснение: Каждая пара в массиве является хорошей.

Пример 3:

Input: nums = [1,2,3]

Output: 0

Ограничения:

1 <= nums.length <= 100

1 <= nums [i] <= 100

func numIdenticalPairs(nums []int) int {

}
https://leetcode.com/problems/number-of-good-pairs/
*/

package main

func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}
