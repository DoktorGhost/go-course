/*
8. K-й пропущенный положительный номер
Дан массив arr положительных целых чисел, отсортированных в строго возрастающем порядке, и целое число k.

Верните k-ое положительное целое число, которое отсутствует в этом массиве.

Пример 1:

Input: arr = [2,3,4,7,11], k = 5

Output: 9

Объяснение:

Отсутствующие положительные целые числа: [1,5,6,8,9,10,12,13, ...]. Пятое отсутствующее положительное целое число - 9.

Пример 2:

Input: arr = [1,2,3,4], k = 2

Output: 6

Объяснение:

Отсутствующие положительные целые числа - [5,6,7,...]. Второе пропущенное положительное целое число - 6.

Ограничения:

1 <= arr.length <= 1000

1 <= arr[i] <= 1000

1 <= k <= 1000

arr[i] < arr[j] for 1 <= i < j <= arr.length

func findKthPositive(arr []int, k int) int {

}
https://leetcode.com/problems/kth-missing-positive-number/
*/
package main

func findKthPositive(arr []int, k int) int {
	res := 1

	count := 0
	idx := 0

	for i := 1; count < k; i++ {
		if idx <= len(arr)-1 {
			if i == arr[idx] {
				idx++
			} else {
				res = i
				count++
			}
		} else {
			res = i
			count++
		}

	}
	return res
}
