/*
4. Построить массив из перестановки
Дана нулевая перестановка nums (с нулевым индексом), создайте массив ans одинаковой длины, где ans[i] = nums[nums[i]] для каждого 0 <= i < nums.length и верните его. Нулевая перестановка nums - это массив уникальных целых чисел от 0 до nums.length - 1 (включительно).

Пример 1:

Input: nums = [0,2,1,5,3,4]

Output: [0,1,2,4,5,3]

Объяснение: Массив ans строится следующим образом: ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]] = [nums[0], nums[2], nums[1], nums[5], nums[3], nums[4]] = [0,1,2,4,5,3]

Пример 2:

Input: nums = [5,0,1,2,3,4]

Output: [4,5,0,1,2,3]

Объяснение: Массив ans строится следующим образом: ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]] = [nums[5], nums[0], nums[1], nums[2], nums[3], nums[4]] = [4,5,0,1,2,3]

Ограничения:

• 1 <= nums.length <= 1000

• 0 <= nums[i] < nums.length

• Элементы в nums уникальны.

func buildArray(nums []int) []int {

}
https://leetcode.com/problems/build-array-from-permutation/

*/

package main

func buildArray(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	for i, v := range nums {
		result[i] = nums[v]
	}
	return result
}
