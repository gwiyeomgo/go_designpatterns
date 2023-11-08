
<p align="center">
  <h1 align="center">
    Go Patterns
  </h1>
</p>

go 언어를 사용한  디자인 패턴 코드를 찾거나 예제 코드를 만들어 모아놓았습니다.

## Creational Patterns

|      |                                                                Pattern | Description                          | Status |
|------|-----------------------------------------------------------------------:|:-------------------------------------|:------:|
| 1    | [Abstract Factory](https://github.com/gwiyeomgo/go_designpatterns/blob/main/01_creational_patterns/abstractFactory/README.md) |  |✔|
| 2    |                  [Builder](https://github.com/gwiyeomgo/go_designpatterns/blob/main/01_creational_patterns/builder/README.md) |                                      | ✔ |
| 3    |     [Factory Method](https://github.com/gwiyeomgo/go_designpatterns/blob/main/01_creational_patterns/factoryMethod/README.md) |                                      | ✔ |
| 4    |              [Prototype](https://github.com/gwiyeomgo/go_designpatterns/blob/main/01_creational_patterns/prototype/README.md) |                                     | ✔ |
| 5    |                [Singleton](https://github.com/gwiyeomgo/go_designpatterns/blob/main/01_creational_patterns/builder/README.md) |                                      | ✔ |

## Structural Patterns

|     |                          Pattern                          | Description | Status |
|-----|:---------------------------------------------------------:|:----------- |:------:|
| 6   |   [Adapter](https://github.com/gwiyeomgo/go_designpatterns/blob/main/02_structural_patterns/adapter/README.md)   | | ✔ |
| 7   |    [Bridge](https://github.com/gwiyeomgo/go_designpatterns/blob/main/02_structural_patterns/bridge/README.md)    | | ✔ |
| 8   | [Composite](https://github.com/gwiyeomgo/go_designpatterns/blob/main/02_structural_patterns/composite/README.md) |  | ✔ |
| 9   | [Decorator](https://github.com/gwiyeomgo/go_designpatterns/blob/main./02_structural_patterns/decorator/README.md) |  | ✔ |
| 10  |    [Facade](https://github.com/gwiyeomgo/go_designpatterns/blob/main/02_structural_patterns/facade/README.md)    |  | ✔ |
| 11  | [Flyweight](https://github.com/gwiyeomgo/go_designpatterns/blob/main/02_structural_patterns/flyweight/README.md) |  | ✔ |
| 12  |     [Proxy](https://github.com/gwiyeomgo/go_designpatterns/blob/main/02_structural_patterns/proxy/README.md)     |  | ✔ |


## Behavioral Patterns

|     |                                         Pattern                                         | Description | Status |
|-----|:---------------------------------------------------------------------------------------:|:----------- |:------:|
| 13  | [ChainOfResponsibilities](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/chain_of_responsibilities/README.md) | | ✔ |
| 14  |                  [Command](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/command/README.md)                  | | ✔ |
| 15  |              [Interpreter](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/interpreter/README.md)              | | ✔ |
| 16  |                 [Iterator](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/iterator/README.md)                 | | ✔ |
| 17  |                 [Mediator](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/mediator/README.md)                 | | ✔ |
| 18  |                  [Memento](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/memento/README.md)                  | | ✔ |
| 19  |                 [Observer](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/observer/README.md)                 | | ✔ |
| 20  |                    [State](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/state/README.md)                    | | ✔ |
| 21  |                 [Strategy](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/strategy/README.md)                 | | ✔ |
| 22  |           [TemplateMethod](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/templateMethod/README.md)           | | ✔ |
| 23  |                  [Visitor](https://github.com/gwiyeomgo/go_designpatterns/blob/main/03_behavioral_patterns/visitor/README.md)                  | | ✔ |

## Microservice design patterns
|         | Pattern                                         | Description | Status |
|---------|-------------------------------------------------|---|---|
| 1       | Service Registry                                |||
| 2       | Circuit Breaker                                 |||
| 3       | API Gateway                                     |||
| 4       | Event-Driven Architecture                       |||
| 5       | Database per Service                            |||
| 6       | Command Query Responsibility Segregation (CQRS) |||
| 7       | Externalized Configuration                      |||
| 8       | Saga Pattern                                    |||
| 9       | Bulkhead Pattern                                |||
| 10      | Backends for Frontends (BFF)                    |||
| 11 (추가) | Retry                                           |https://engineering.mercari.com/en/blog/entry/20210126-retry-pattern-in-microservices/||
| 12 (추가)| Circuit Debounce                                |||

## Go Concurrency Patterns
https://go.dev/blog/pipelines
https://woojinger.tistory.com/82

|         | Pattern | Description | Status |
|---------|---------|---|---|
| 1       | fan-in  |||
| 2       | fan-out |||
