/*
7. Защита IP-адреса
Дан действительный (IPv4) IP-адрес, вернуть защищенную версию этого IP-адреса.

Защищенный IP-адрес заменяет каждую точку "." с "[.]"

Пример 1:

Input: адрес = "1.1.1.1"

Output: "1[.]1[.]1[.]1"

Пример 2:

Input: адрес = "255.100.50.0"

Output: "255[.]100[.]50[.]0"

Ограничения:

Данный адрес является действительным адресом IPv4.

func defangIPaddr(address string) string {

}
https://leetcode.com/problems/defanging-an-ip-address/
*/

package main

import "strings"

func defangIPaddr(address string) string {
	result := strings.ReplaceAll(address, ".", "[.]")
	return result
}
