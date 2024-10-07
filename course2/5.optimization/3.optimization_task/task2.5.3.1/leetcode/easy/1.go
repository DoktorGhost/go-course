/*
1. N-th Tribonacci Number
Последовательность Трибоначчи Tn определяется следующим образом:

T0 = 0, T1 = 1, T2 = 1 и Tn + 3 = Tn + Tn + 1 + Tn + 2 для n> = 0.

Дано n, верните значение Tn.

Пример 1:

Input: n = 4

Output: 4

Объяснение:

T_3 = 0 + 1 + 1 = 2

T_4 = 1 + 1 + 2 = 4

Пример 2:

Input: n = 25

Output: 1389537

Ограничения:

0 <= n <= 37
Ответ гарантированно поместится в 32-битное целое число, т.е. ответ <= 2 ^ 31 - 1.
func tribonacci(n int) int {

}

https://leetcode.com/problems/n-th-tribonacci-number/submissions/1392961851/
*/
package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	n0 := 0
	n1 := 1
	n2 := 1

	for i := 3; i <= n; i++ {
		n3 := n0 + n1 + n2
		n0 = n1
		n1 = n2
		n2 = n3
	}

	return n2
}
