package main

func main() {
	mailFacade := newMailFacade("sdfsdfew",
		"t@gmail.com",
		[]string{"t@gmail.com"},
		[]byte("Hello World"),
	)
	mailFacade.Send()
}
