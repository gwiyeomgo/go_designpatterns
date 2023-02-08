package main

// 요청을 처리하는 쪽과 디커플링 되어있다.
// 클라이언트 코드가 어떤 행동을 하는지 관여하지 않음
// 클라이언트에서 체인만 구성
func main() {

	printMessage := &Print{}

	//Set next for medical department
	logging := &Logging{}
	logging.setNext(printMessage)

	//Set next for reception department
	auth := &Auth{}
	auth.setNext(logging)

	request := &Request{message: "요청합니다~~~"}
	//Patient visiting
	auth.execute(request)
}
