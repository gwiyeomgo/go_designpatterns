package case2

/*

from := "you@gmail.com"
auth := smtp.PlainAuth("", from, "<your gmail password>", "smtp.gmail.com")
err := smtp.SendMail(
    "smtp.gmail.com:587",               // server address
    auth,                               // authentication
    from,                               // sender's address
    []string{"recipient@example.com"},  // recipients' address
    []byte("Hello World!"),             // message body
)
if err != nil {
    log.Print(err)
}

*/

func main() {
	settings := Settings{}
	sender := NewEmailSender(settings)
	body := "Hello World"
	err := sender.Send([]string{"me@example.com"}, []byte(body))
	if err != nil {
		//log.Errorf("unexpected error: %s", err)
	}

}
