/*
14. Богатство самого богатого клиента
Вам дана матрица целых чисел accounts, где accounts[i][j] - это сумма денег, которые имеет i-й клиент в j-м банке. Верните богатство самого богатого клиента.

Богатство клиента - это сумма денег, которые он имеет во всех своих банковских счетах. Самым богатым клиентом считается клиент с максимальным богатством.

Пример 1:

Input: accounts = [[1,2,3],[3,2,1]]

Output: 6

Объяснение:

1-й клиент имеет богатство = 1 + 2 + 3 = 6

2-й клиент имеет богатство = 3 + 2 + 1 = 6

Оба клиента считаются самыми богатыми с богатством 6 каждый, поэтому возвращаем 6.

Пример 2:

Input: accounts = [[1,5],[7,3],[3,5]]

Output: 10

Объяснение:

1-й клиент имеет богатство = 6

2-й клиент имеет богатство = 10

3-й клиент имеет богатство = 8

2-й клиент является самым богатым с богатством 10.

Пример 3:

Input: accounts = [[2,8,7],[7,1,3],[1,9,5]]

Output: 17

Ограничения:

m == accounts.length

n == accounts[i].length

1 <= m, n <= 50

1 <= accounts[i][j] <= 100

func maximumWealth(accounts [][]int) int {

}
https://leetcode.com/problems/richest-customer-wealth/
*/

package main

func maximumWealth(accounts [][]int) int {
	maxx := 0

	for _, account := range accounts {
		promMax := 0
		for _, val := range account {
			promMax = promMax + val
		}
		if promMax > maxx {
			maxx = promMax
		}
	}
	return maxx
}
