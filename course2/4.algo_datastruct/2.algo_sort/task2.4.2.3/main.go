package main

// Структура пользователя
type User struct {
	ID   int
	Name string
	Age  int
}

// Функция слияния двух отсортированных массивов пользователей
func Merge(arr1 []User, arr2 []User) []User {
	if len(arr1) < 1 {
		return arr2
	}
	if len(arr2) < 1 {
		return arr1
	}
	mergedArr := []User{}

	i := 0
	j := 0
	for {
		if arr1[i].ID < arr2[j].ID {
			mergedArr = append(mergedArr, arr1[i])
			i++
		} else {
			mergedArr = append(mergedArr, arr2[j])
			j++
		}

		if j == len(arr2) {
			mergedArr = append(mergedArr, arr1[i:]...)
			return mergedArr
		}
		if i == len(arr1) {
			mergedArr = append(mergedArr, arr2[j:]...)
			return mergedArr
		}
	}
}
