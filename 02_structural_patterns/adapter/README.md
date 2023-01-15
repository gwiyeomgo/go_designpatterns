기존 코드를 클라인언트가 사용하는 인터페이스의 구현체로 바꿔주는 패턴
* 클라이언트가 사용하는 인터페이스를 따르지 않는 기존 코드를 재사용할 수 있게 해준다
```
클라이언드 = 타겟인터페이스

어뎁터 = 둘사이 연결 어뎁터

어뎁티 = 한국형 냉장고
```

장점
* 기존 코드를 변경하지 않고 새로운 유형의 어댑터를 도입할 수 있습니다.
* 기존 코드가 하던 일과 특정 인터페이스 구현체로 변환하는 작업을 
각기 다른 클래스로 분리하여 관리할 수 있다

단점
* 새 인터페이스와 구조를 도입해야 하므로 복잡도가 증가할 수 있다.

ex) https://github.com/gwiyeomgo/nomadcoin/commit/ae60a38da901b3ecfccd2224247e33c39f7f7f6f

https://up-to-date-items.tistory.com/226
https://medium.com/@josueparra2892/adapter-pattern-in-go-d77e08abd526
https://refactoring.guru/design-patterns/adapter/go/example#example-0