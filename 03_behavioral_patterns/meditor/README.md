# 이터레이터
집합 객체가 갖고 있는 객체들에 손쉽게 접근할 수 있다
집합 객체가 어떠한 구조로 이루어 졌는지 알 필요 없다
힙인지 스택인지 리스트인지 알 필요없이
클라이언트는 hasNext getNext 만 갖고 순화할 수 있다
순회하는 방법만 안다
구체적인 순회로직은 안쪽에 존재함

클라이언트 입장에서는
기존에 큰 변경사항 없이

콘크리트 에그리에이트 입장에서는 내용을 변경한다
새 이러터이터를 만들면 이터레이터를 호출하는 코드를 추가해줘야 하기도 함

단점은 이터레이터를 만드는 것이 유용한 것인지 판단이 필요

# 중재자 패턴 (mediator)
여러 객체들간의 소통한는 방법을 추상화 시켜서 커플링을 낮춘다

ex) 층간 소음 등으로 아파트 공지(관리사부소에서)
ex) 비행기 관제탑이 비행기들간 연락을 알림
ex) 호텔의 클리닉 서비스


여러 컴포넌트간의 결합도를 중재자를 통해 낮출 수 있다

모즌 colleague 들이 mediator 를 참조하고 있다
colleague 가 colleague 를 참조하지 않는다
classes do not communicate directly with each other.


https://dev.to/tomassirio/mediator-design-pattern-in-go-15ma
https://medium.com/@josueparra2892/mediator-pattern-in-go-be0769ef0a45
https://pavelfokin.dev/blog/mediator-pattern-in-go/