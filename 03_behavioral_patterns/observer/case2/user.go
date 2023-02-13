package main

import "fmt"

// 구현체 (subscriber 의)
type User struct {
	name string
}

func newUser(name string) *User {
	return &User{
		name: name,
	}
}
func (u *User) getName() string {
	return u.name
}
func (u *User) handleMessage(message string) {
	fmt.Println(message)
}
