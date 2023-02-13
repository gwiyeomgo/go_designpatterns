package main

type ChatServer struct {
	subscribers map[string][]Subscriber //주제별로 subscriber 의 목록을 관리
}

func (i *ChatServer) register(subject string, o Subscriber) {
	subscribers := i.subscribers[subject]
	if len(subscribers) > 0 {
		subscribers = append(subscribers, o)
		i.subscribers[subject] = subscribers
	} else {
		list := make([]Subscriber, 0)
		list = append(list, o)
		if i.subscribers == nil {
			i.subscribers = map[string][]Subscriber{
				subject: list,
			}
		} else {
			i.subscribers[subject] = list
		}

	}
}

func (i *ChatServer) unregister(subject string, o Subscriber) {
	subscribers := i.subscribers[subject]
	observerListLength := len(subscribers)
	if len(subscribers) > 0 {
		for j, observer := range subscribers {
			if observer == o {
				subscribers[observerListLength-1], subscribers[j] = subscribers[j], subscribers[observerListLength-1]
				i.subscribers[subject] = subscribers[:observerListLength-1]
			}
		}
	}
}

func (i *ChatServer) sendMessage(user *User, subject string, message string) {
	subscribers := i.subscribers[subject]
	if len(subscribers) > 0 {
		userMessage := user.getName() + ": " + message
		for _, subscriber := range subscribers {
			subscriber.handleMessage(userMessage)
		}

	}
}
