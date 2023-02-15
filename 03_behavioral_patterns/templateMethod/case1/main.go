package main

import "fmt"

func main() {
	smsOTP := &Sms{}
	//상위 클래스에 하위 클래스 넣음
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(6)
}
