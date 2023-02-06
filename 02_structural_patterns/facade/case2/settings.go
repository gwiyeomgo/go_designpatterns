package main

type Settings struct {
	ServerHost string
	ServerPort string
	Password   string
}

func newSettings(Password string) *Settings {
	return &Settings{
		ServerHost: "smtp.gmail.com",
		ServerPort: "587",
		Password:   Password,
	}
}
