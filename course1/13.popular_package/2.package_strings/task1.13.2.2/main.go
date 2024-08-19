package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

func main() {
	users := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
				{Message: "Bad Comment 4"},
			},
		},
	}

	users = FilterComments(users)
	fmt.Println(users)

}

func FilterComments(users []User) []User {
	resUsers := make([]User, len(users), len(users))
	for i, user := range users {
		resUsers[i].Name = user.Name
		resUsers[i].Comments = GetGoodComments(user)
	}
	return resUsers
}

func IsBadComment(comment string) bool {
	//содержит ли комментарий "bad comment"
	res := strings.ToLower(comment)
	return strings.Contains(res, "bad comment")

}

func GetBadComments(user User) []Comment {
	result := []Comment{}

	for _, comment := range user.Comments {
		if IsBadComment(comment.Message) {
			result = append(result, comment)
		}
	}

	return result
}

func GetGoodComments(user User) []Comment {
	result := []Comment{}

	for _, comment := range user.Comments {
		if !IsBadComment(comment.Message) {
			result = append(result, comment)
		}
	}

	return result
}
