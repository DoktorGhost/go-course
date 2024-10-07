/*
5. Максимальная сумма двух звеньев связанного списка
В связанном списке размера n, где n является четным, i-й узел (с 0 индексом) связанного списка известен как близнец(twin) (n-1-i) узла, если 0 <= i <= (n / 2) - 1.

Например, если n = 4, то узел 0 является близнецом узла 3, а узел 1 является близнецом узла 2. Это единственные узлы с близнецами для n = 4.
Сумма близнецов(twin sum) определяется как сумма узла и его близнеца.

Дан заголовок связанного списка с четной длиной, верните максимальную сумму близнецов связанного списка.

Пример 1:

Input: head = [5,4,2,1]

Output: 6

Объяснение:

Узлы 0 и 1 являются близнецами узлов 3 и 2 соответственно. Все имеют сумму близнецов = 6. В связанном списке нет других узлов с близнецами. Таким образом, максимальная сумма близнецов в связанном списке составляет 6.

Пример 2:

Input: head = [4,2,2,3]

Output: 7

Объяснение:

Узлы с двойниками в этом связанном списке:

- Узел 0 является двойником узла 3 с двойной суммой 4 + 3 = 7.

- Узел 1 является двойником узла 2 с двойной суммой 2 + 2 = 4.

Таким образом, максимальная двойная сумма связанного списка равна max(7, 4) = 7.

Пример 3:

Input: head = [1,100000]

Output: 100001

Объяснение:

В связанном списке с двойным суммой 1 + 100000 = 100001 есть только один узел с двойником.

Constraints:

Число узлов в списке является четным целым числом в диапазоне [2, 105].
1 <= Node.val <= 105
https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
*/

package main

func pairSum(head *ListNode) int {
	node := head
	result := make(map[int]int)
	maxx := 0
	i := 0
	for node.Next != nil {
		result[i] = node.Val
		i++
		node = node.Next
	}
	result[i] = node.Val

	n := len(result) - 1

	for i := 0; i <= n/2; i++ {
		a := result[i] + result[n-i]
		if maxx < a {
			maxx = a
		}
	}

	return maxx
}