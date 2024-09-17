/*
2. Объединение массива
Дан массив целых чисел nums длиной n. Вы должны создать массив ans длиной 2n, где ans[i] == nums[i] и ans[i + n] == nums[i] для 0 <= i < n (с индексацией 0).

ans является слиянием двух массивов nums.

Верните массив ans.

Пример 1:

Input: nums = [1,2,1]

Output: [1,2,1,1,2,1]

Объяснение: Массив ans формируется следующим образом:

- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]

- ans = [1,2,1,1,2,1]

Пример 2:

Input: nums = [1,3,2,1]

Output: [1,3,2,1,1,3,2,1]

Объяснение: Массив ans формируется следующим образом:

- ans = [nums[0],nums[1],nums[2],nums[3],nums[0],nums[1],nums[2],nums[3]]

- ans = [1,3,2,1,1,3,2,1]

Constraints:

n == nums.length
1 <= n <= 1000
1 <= nums[i] <= 1000
func getConcatenation(nums []int) []int {

}
https://leetcode.com/problems/concatenation-of-array/
*/

package main

func getConcatenation(nums []int) []int {
	result := make([]int, len(nums)*2, len(nums)*2)
	for i := 0; i < len(nums)*2; i++ {
		j := i
		if i >= len(nums) {
			j = i - len(nums)
		}
		result[i] = nums[j]
	}

	return result
}
