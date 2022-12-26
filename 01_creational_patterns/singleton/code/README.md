

Lazy Initialization: :게으른 초기화 vs eager initialization


> golang 을 사용해서 고루틴을 사용해서  LazyInit 를  여러번 호출할 경우
race condition 을 막시 위해 여러 방법이 존재한다

> race condition ? 2개 스레드가 하나의 자원을 놓고 경쟁하는 상황[출처](https://iredays.tistory.com/125)
> 1. Lock 을 통해  race condition  막기
> 2. Sync.Once 를 사용

하지만  이중 확인 잠금을 사용하면 지연 초기화를 멱등적이고 스레드로부터 안전하게 만들 수 있습니다.
이론적으로는 그럴듯해 보이지만 실제 응용 프로그램에서 실제 데이터가 확인 될 때 까지 초기화 연기 된다


---
tip
>If you're a Java developer you might recognize this as a safe approach to double-checked locking.
In Java, the volatile keyword is typically used on instance instead of using a read/write lock,
but since Go does not have a volatile keyword
(there is sync.atomic, and we'll get to that) we've gone with a read lock.
[출처](https://launchdarkly.com/blog/golang-pearl-thread-safe-writes-and-double-checked-locking-in-go/)


---
tip
>Package init() functions are guaranteed to be called only once
and all called from a single thread
( they're thread-safe unless you make them multi-threaded)
[출처](https://medium.com/@ishagirdhar/singleton-pattern-in-golang-9f60d7fdab23)

---
reflection ?

---

직력화 역직렬화?

>직렬화 (Serialization) 란 객체의 상태를 보관이나 전송 가능한 상태로 변환하는 것입니다. 구조체 만이 아니라 숫자, 문자열, 배열, 맵 역시 직렬화가 가능합니다. 직렬화의 반대로 보관되거나 전송받은 것을 다시 객체로 복원하는 것을 역직렬화 (Deserialization) 라 합니다.
보조기억장치에 저장 및 불러오기 / 네트워크를 통한 메시지 전송 / RPC 등의 방법을 사용할 때 직렬화가 필요합니다.