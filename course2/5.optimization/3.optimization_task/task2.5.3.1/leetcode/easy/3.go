/*
3. Преобразование температуры
Вам дается неотрицательное дробное число с двумя знаками после запятой, указывающее температуру в градусах Цельсия.

Вы должны преобразовать Цельсий в Кельвины и Фаренгейты и вернуть его в виде массива ans = [kelvin, fahrenheit].

Верните массив ans. Ответы в пределах 10-5 от реального ответа будут приняты.

Обратите внимание:

Kelvin = Celsius + 273.15

Fahrenheit = Celsius * 1.80 + 32.00

Пример 1:

Input: celsius = 36.50

Output: [309.65000,97.70000]

Объяснение: Температура при 36,50 Цельсия переведена в Кельвины равна 309,65, а в Фаренгейты - 97,70.

Пример 2:

Input: celsius = 122.11

Output: [395.26000,251.79800]

Объяснение: Температура при 122,11 Цельсия преобразована в Кельвинах составляет 395,26, а в Фаренгейтах - 251,798.

Constraints:

0 <= celsius <= 1000
func convertTemperature(celsius float64) []float64 {

}
https://leetcode.com/problems/convert-the-temperature/

*/

package main

func convertTemperature(celsius float64) []float64 {
	result := make([]float64, 2, 2)
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00
	result[0] = kelvin
	result[1] = fahrenheit

	return result
}
