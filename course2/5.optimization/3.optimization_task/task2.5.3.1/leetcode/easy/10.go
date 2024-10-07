/*
10. Перетасовать массив
Дан массив nums, состоящий из 2n элементов в виде [x1,x2,...,xn,y1,y2,...,yn].

Верните массив в виде [x1,y1,x2,y2,...,xn,yn].

Пример 1:

Input: nums = [2,5,1,3,4,7], n = 3

Output: [2,3,5,4,1,7]

Объяснение:

Поскольку x1=2, x2=5, x3=1, y1=3, y2=4, y3=7, то ответ равен [2,3,5,4,1,7].

Пример 2:

Input: nums = [1,2,3,4,4,3,2,1], n = 4

Output: [1,4,2,3,3,2,4,1]

Пример 3:

Input: nums = [1,1,2,2], n = 2

Output: [1,2,1,2]

Ограничения:

1 <= n <= 500

nums.length == 2n

1 <= nums[i] <= 10^3

func shuffle(nums []int, n int) []int {

}
https://leetcode.com/problems/shuffle-the-array/
*/

package main

func shuffle(nums []int, n int) []int {
	var result []int
	
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}
