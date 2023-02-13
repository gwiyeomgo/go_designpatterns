package main

func main() {
	user1 := User{"keesun"}
	user2 := User{"witheship"}

	chatServer := ChatServer{}
	chatServer.register("오징어게임", &user1)
	chatServer.register("오징어게임", &user2)
	chatServer.register("디자인패턴", &user1)

	chatServer.sendMessage(&user1, "오징어게임", "아...")
	chatServer.sendMessage(&user2, "디자인패턴", "옵저버 패턴으로 만든 채팅")

	chatServer.unregister("디자인패턴", &user2)

	chatServer.sendMessage(&user2, "디자인패턴", "옵저버 패넡 장단점")
}
