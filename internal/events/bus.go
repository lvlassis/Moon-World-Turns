package events

type Bus struct {
	subscribers map[string][]func(data interface{})
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]func(data interface{})),
	}
}

// Emit emite um evento com payload type-safe
func Emit[T any](b *Bus, eventName string, data T) {
	if callbacks, found := b.subscribers[eventName]; found {
		for _, callback := range callbacks {
			callback(data)
		}
	}
}

// On registra um listener type-safe para um evento
func On[T any](b *Bus, eventName string, callback func(T)) {
	b.subscribers[eventName] = append(b.subscribers[eventName], func(data interface{}) {
		if payload, ok := data.(T); ok {
			callback(payload)
		}
	})
}

