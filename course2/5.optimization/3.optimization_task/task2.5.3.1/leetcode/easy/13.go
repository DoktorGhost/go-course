/*
13. Драгоценности и камни
Вам даны строки jewels, представляющие типы камней, которые являются драгоценностями, и камни, представляющие stones, которые у вас есть. Каждый символ в stones является типом камня, который у вас есть. Вы хотите знать, сколько из камней, которые у вас есть, также являются драгоценностями.

Буквы чувствительны к регистру, поэтому «a» считается разным типом камня от «A».

Пример 1:

Input: jewels = "aA", stones = "aAAbbbb"

Output: 3

Пример 2:

Input: jewels = "z", stones = "ZZ"

Output: 0

Ограничения:

1 <= jewels.length, stones.length <= 50

jewels и stones состоят только из английских букв.

Все символы jewels уникальны.

func numJewelsInStones(jewels string, stones string) int {

}
https://leetcode.com/problems/jewels-and-stones/
*/

package main

func numJewelsInStones(jewels string, stones string) int {
	maps := make(map[rune]int)
	res := 0
	jewelsRune := []rune(jewels)
	stonesRune := []rune(stones)
	for _, i := range stonesRune {
		maps[i]++
	}
	for _, v := range jewelsRune {
		res = res + maps[v]
	}
	return res
}
