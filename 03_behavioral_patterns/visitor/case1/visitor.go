package main

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}

//java 에서는 visitor의 메서드들이
//이름은 같으나 파라미터가 다르게  메서드 오버로딩 할 수 있음
//golang 은 인터페이스 내 이름이 같으면 중복으로 인식
