/*
9. Арифметические подмассивы
Последовательность чисел называется арифметической, если она состоит как минимум из двух элементов, а разница между каждыми двумя соседними элементами одинакова. Более формально, последовательность s является арифметической, если и только если s[i + 1] - s[i] == s[1] - s[0] для всех допустимых i.

Например, это арифметические последовательности:

Следующая последовательность не является арифметической:

Вам дан массив из n целых чисел, nums и два массива из m целых чисел каждый, l и r, представляющий m диапазонов запросов, где i-й запрос является диапазоном [l [i], r [i]]. Все массивы имеют индекс 0.

Верните список булевых элементов ответа, где answer [i] является true, если подмассив nums [l [i]], nums [l [i] + 1], ..., nums [r [i]] можно переупорядочить для формирования арифметической последовательности, и false в противном случае.

Пример 1:

Input: nums = [4,6,5,9,3,7], l = [0,0,2], r = [2,3,5]

Output: [true,false,true]

Объяснение:

В 0-м запросе подмассив состоит из [4,6,5]. Это можно переставить как [6,5,4], что является арифметической прогрессией.

В 1-м запросе подмассив состоит из [4,6,5,9]. Это не может быть переставлено в арифметическую прогрессию.

Во 2-м запросе подмассив состоит из [5,9,3,7]. Это может быть переставлено как [3,5,7,9], что является арифметической прогрессией.Arithmetic Subarrays

Пример 2:

Input: nums = [-12,-9,-3,-12,-6,15,20,-25,-20,-15,-10], l = [0,1,6,4,8,7], r = [4,4,9,7,9,10]

Output: [false,true,false,false,true,true]

Constraints:

n == nums.length
m == l.length
m == r.length
2 <= n <= 500
1 <= m <= 500
0 <= l[i] < r[i] < n
-105 <= nums[i] <= 105
https://leetcode.com/problems/arithmetic-subarrays/
*/

package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, 0)
	for i := 0; i < len(l); i++ {
		copied := make([]int, len(nums))
		copy(copied, nums)
		temp := copied[l[i] : r[i]+1]
		if len(temp) < 2 {
			ans = append(ans, false)
			continue
		}
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})

		delta := temp[1] - temp[0]
		for j := 1; j < len(temp)-1; j++ {
			if temp[j+1]-temp[j] != delta {
				ans = append(ans, false)
				break
			}
		}
		if len(ans)-1 < i {
			ans = append(ans, true)
		}
	}
	return ans
}
