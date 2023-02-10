
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