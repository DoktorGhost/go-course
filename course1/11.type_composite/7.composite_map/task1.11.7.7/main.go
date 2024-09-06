package main

import "fmt"

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	var workUsersArr []User

	uniqMap := make(map[string]struct{})

	for _, user := range users {
		if _, ok := uniqMap[user.Nickname]; !ok {
			uniqMap[user.Nickname] = struct{}{}
			workUsersArr = append(workUsersArr, user)
		}
	}

	result := make([]User, len(workUsersArr))
	copy(result, workUsersArr)
	return result
}

func main() {
	users := []User{{Nickname: "John", Age: 25, Email: "john@gmail.com"}, {Nickname: "Jane", Age: 25, Email: "jane@gmail.com"}, {Nickname: "John2", Age: 48, Email: "john24@gmail.com"}}
	uniq := getUniqueUsers(users)

	fmt.Println(uniq)

}
