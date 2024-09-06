package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Word struct {
	Word string
	Pos  int
}

func filterWords(text string, censorMap map[string]string) string {
	sentences := splitSentences(text)

	var result []string

	// Обработка каждого предложения
	for _, sentence := range sentences {
		words := strings.Fields(sentence)
		uniqueWords := make(map[string]Word) // Создание пустой карты уникальных слов
		var newWords []string

		for _, word := range words {
			// Проверка, содержит ли карта цензурных слов текущее слово
			if replacement, found := censorMap[strings.ToLower(word)]; found {
				// Замена слова на значение из карты, используя CheckUpper
				newWords = append(newWords, CheckUpper(word, replacement))
			} else {
				// Проверка ключа в карте уникальных слов
				if _, exists := uniqueWords[strings.ToLower(word)]; !exists {
					uniqueWords[strings.ToLower(word)] = Word{Word: word, Pos: len(newWords)}
					newWords = append(newWords, word)
				}
			}
		}

		// Возвращение предложения из слайса слов
		result = append(result, WordsToSentence(newWords))
	}

	return strings.Join(result, " ")
}

// удаляет пустые слова из слайса и объединяет их предложение, добавляя в конце "!"
func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))
	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}
	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

// проверяет, нужно ли заменять первую букву на заглавную
func CheckUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}
	chars := []rune(old)

	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
	}
	return new
}

// разделяет текст на предложения
func splitSentences(message string) []string {
	originSentences := strings.Split(message, "!")

	var orphan string
	var sentences []string

	for i, sentence := range originSentences {
		sentence = strings.TrimSpace(sentence)
		words := strings.Split(sentence, " ")
		if len(words) == 1 {
			if len(orphan) > 0 {
				orphan += " "
			}
			orphan += words[0] + "!"
			continue
		}
		if orphan != "" {
			originSentences[i] = strings.Join([]string{orphan, " ", sentence}, " ") + "!"
			orphan = ""
		}
		sentences = append(sentences, originSentences[i])
	}
	return sentences
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!"

	censorMap := map[string]string{
		"крипта":   "фрукты",
		"крипту":   "фрукты",
		"крипты":   "фруктов",
		"биткоин":  "фрукты",
		"лфйткоин": "фрукты",
		"эфир":     "фрукты",
	}

	filteredText := filterWords(text, censorMap)
	fmt.Println(filteredText)

}
