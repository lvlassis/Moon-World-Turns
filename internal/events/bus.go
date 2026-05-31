package events

type Bus struct {
	subscribers map[string][]func(evname string, ev interface{})
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]func(evname string, ev interface{})),
	}
}

func (b *Bus) Subscribe(eventName string, callback func(evname string, ev interface{})) {
	b.subscribers[eventName] = append(b.subscribers[eventName], callback)
}

func (b *Bus) Publish(eventName string, eventData interface{}) {
	if callbacks, found := b.subscribers[eventName]; found {
		for _, callback := range callbacks {
			callback(eventName, eventData)
		}
	}
}

