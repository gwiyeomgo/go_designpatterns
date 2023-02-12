메멘토 패턴은 행동 디자인 패턴입니다. 이 패턴은 객체 상태의 스냅숏을 만든 후 나중에 복원할 수 있도록 합니다

스냅샷

장점
originator (Game)의 상태를 저장해 
caretaker 를 사용한다면 단일책임의 원칙


단점
메멘토 객체가 많은 정보를 갖고 있고
자주 생성한다면 메모리 사용량에 많은 영향을 줄 수 있다
오래된 메멘토를 정리하는 역할도 caretaker 가 할수도 있다

https://refactoring.guru/ko/design-patterns/memento/go/example