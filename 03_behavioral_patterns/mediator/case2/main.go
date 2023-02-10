package main

import "fmt"

type User struct {
	Name     string
	Messages []string
}

// 중재자
type Chat struct {
	Users map[string]*User
}

func NewChat() *Chat {
	return &Chat{make(map[string]*User)}
}

func (c *Chat) Add(user User) {
	c.Users[user.Name] = &user
}

func (c *Chat) Say(to User, msg string) error {
	user, ok := c.Users[to.Name]
	if !ok {
		return fmt.Errorf("%s not in the chat\n", to.Name)
	}
	user.Messages = append(user.Messages, msg)
	return nil
}

func (c *Chat) SayAll(msg string) {
	for _, user := range c.Users {
		user.Messages = append(user.Messages, msg)
	}
}

func main() {
	// Create users (최초 가입)
	John := User{Name: "John"}
	Bob := User{Name: "Bob"}
	Alice := User{Name: "Alice"}
	// Create chat
	chat := NewChat()
	// Add users to the chat
	chat.Add(Bob)
	chat.Add(John)
	chat.Add(Alice)
	// Send messages
	chat.Say(Bob, "Hello, Bob!")
	chat.Say(Alice, "Hello, Alice!")
	chat.SayAll("Hello, All!")

	bob := chat.Users[Bob.Name]
	fmt.Println(bob.Messages)

	john := chat.Users[John.Name]
	fmt.Println(john.Messages)

	alice := chat.Users[Alice.Name]
	fmt.Println(alice.Messages)
}
