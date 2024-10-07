/*
11. Текущее значение суммы в одномерном массиве
Дан массив nums. Мы определяем текущее значение суммы массива как runningSum[i] = sum(nums [0]…nums [i]).

Верните текущее значение суммы nums.

Пример 1:

Input: nums = [1,2,3,4]

Output: [1,3,6,10]

Объяснение: Текущее значение суммы получается следующим образом: [1, 1+2, 1+2+3, 1+2+3+4].

Пример 2:

Input: nums = [1,1,1,1,1]

Output: [1,2,3,4,5]

Объяснение: Текущее значение суммы получается следующим образом: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

Пример 3:

Input: nums = [3,1,2,10,1]

Output: [3,4,6,16,17]

Ограничения:

1 <= nums.length <= 1000

-10^6 <= nums [i] <= 10^6

func runningSum(nums []int) []int {

}
https://leetcode.com/problems/running-sum-of-1d-array/
*/

package main

func runningSum(nums []int) []int {
	var res []int
	count := 0

	for _, v := range nums {
		res = append(res, v+count)
		count = count + v
	}
	return res
}
