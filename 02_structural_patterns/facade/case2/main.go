package main

import "log"

func main() {
	mailFacade := newMailFacade("sdfsdfew",
		"t@gmail.com",
		[]string{"t@gmail.com"},
		[]byte("Hello World"),
	)
	err := mailFacade.Send()
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
