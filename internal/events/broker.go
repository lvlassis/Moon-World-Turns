package events

type Broker struct {
	subscribers map[string][]func(evname string, ev interface{})
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]func(evname string, ev interface{})),
	}
}

func (b *Broker) Subscribe(eventName string, callback func(evname string, ev interface{})) {
	b.subscribers[eventName] = append(b.subscribers[eventName], callback)
}

func (b *Broker) Publish(eventName string, eventData interface{}) {
	if callbacks, found := b.subscribers[eventName]; found {
		for _, callback := range callbacks {
			callback(eventName, eventData)
		}
	}
}

