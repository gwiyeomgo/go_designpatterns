
책임들의 연결
단일 책임 원칙에서 말하는 책임
클레스가 변경되어야하는 이유는 1가지 이유여야 한다

클레스들이 연결된 구조로 해결

요청을 보내는 쪽과 처리하는 
디커플링된 상태에서 요청을 처리

request 

요청을 보내는 쪽(sender)과 요청을 처리하는 receiver 의 분리하는 패터
핸들러 체인을 사용해서 요청을 처리


https://refactoring.guru/design-patterns/chain-of-responsibility/go/example#example-0