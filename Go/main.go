package main

import (
	"fmt"
	_ "os"
)

type user struct {
	id   int
	name string
	mail string
	pass string
}

var users []user

var newid = map[int]user{}

func (u *user) rand() {
	for i := range users {
		users[i].id = i + 1
	}
}

func main() {
	andrew := user{
		name: "Andrey",
		mail: "andrey@mail.ru",
		pass: "12334"}
	keril := user{
		name: "Kirill",
		mail: "kirill@mail.ru",
		pass: "123345"}

	users = append(users, andrew)
	users = append(users, keril)
	fmt.Println(users)
	for _, person := range users {
		person.rand()
		fmt.Println(fmt.Sprintf("%s:", person.name), person.id)
		newid[person.id] = person
	}
	fmt.Println(newid)
}
