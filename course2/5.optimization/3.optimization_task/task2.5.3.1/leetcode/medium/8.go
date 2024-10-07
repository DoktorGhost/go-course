/*
8. Минимальное количество вершин для достижения всех узлов
Дан ориентированный ациклический граф, с n вершинами, пронумерованными от 0 до n-1, и массив ребер, где edges [i] = [fromi, toi] представляет собой направленное ребро от узла fromi к узлу toi .

Найдите самый маленький набор вершин, из которых достижимы все узлы в графе. Гарантируется существование уникального решения.

Обратите внимание, что вы можете вернуть вершины в любом порядке.

Пример 1:

Input: n = 6, edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]

Output: [0,3]

Объяснение: Невозможно достичь всех узлов из одной вершины. От 0 мы можем достичь [0,1,2,5]. От 3 мы можем достичь [3,4,2,5]. Так что мы выводим [0,3].].

Пример 2:

Input: n = 5, edges = [[0,1],[2,1],[3,1],[1,4],[2,4]]

Output: [0,2,3]

Объяснение: Обратите внимание, что вершины 0, 3 и 2 недостижимы из любой другой вершины, поэтому мы должны их включить. Также любая из этих вершин может достичь узлов 1 и 4.

Constraints:

2 <= n <= 10^5
1 <= edges.length <= min(10^5, n * (n - 1) / 2)
edges[i].length == 2
0 <= fromi, toi < n
All pairs (fromi, toi) are distinct.
https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
*/

package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	degrees := make([]int, n)
	for _, e := range edges {
		degrees[e[1]]++
	}
	starts := make([]int, 0, n)
	for i, d := range degrees {
		if d == 0 {
			starts = append(starts, i)
		}
	}
	return starts
}
