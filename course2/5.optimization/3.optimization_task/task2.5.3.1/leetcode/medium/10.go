/*
10. Возможности буквенных плиток.
У вас есть n плиток, на каждой из которых напечатана одна буква tiles[i].

Верните количество возможных непустых последовательностей букв, которые вы можете составить, используя буквы, напечатанные на этих плитках.

Пример 1:

Input: tiles = "AAB"

Output: 8

Объяснение : Возможные последовательности: "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".

Пример 2:

Input: tiles = "AAABBC"

Output: 188

Пример 3:

Input: tiles = "V"

Output: 1

Constraints:

1 <= tiles.length <= 7
Набор плиток состоит из заглавных букв английского алфавита..
https://leetcode.com/problems/letter-tile-possibilities/
*/

package main

import "strings"

func numTilePossibilities(tiles string) int {
	seqs := map[string]bool{}
	t := strings.Split(tiles, "")
	n := len(tiles)
	done := make([]bool, n)

	var dfs func(cur string, size int)
	dfs = func(cur string, size int) {
		if len(cur) == size {
			seqs[cur] = true
			return
		}
		for i := 0; i < n; i++ {
			if done[i] == false {
				done[i] = true
				dfs(cur+t[i], size)
				done[i] = false
			}
		}
		return
	}

	for size := 1; size <= n; size++ {
		dfs("", size)
	}
	return len(seqs)
}
