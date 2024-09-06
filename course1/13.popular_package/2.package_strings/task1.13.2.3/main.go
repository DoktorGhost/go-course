package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateRandomString(length int) string {
	var sb strings.Builder
	charset := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnmЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮйцукенгшщзхъфывапролджэячсмитьбюёЁ!№;%:?*()_+=,.&"
	char := []rune(charset)

	for i := 0; i < length; i++ {

		idx := rand.Intn(len(char))
		
		sb.WriteRune(char[idx])
	}

	return sb.String()
}

func main() {
	randomString := GenerateRandomString(10)
	fmt.Println(randomString)
}
