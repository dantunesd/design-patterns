package customer

type EventManager struct {
	listeners []Listener
}

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: []Listener{},
	}
}

func (e *EventManager) AddListener(listener Listener) {
	e.listeners = append(e.listeners, listener)
}

func (e *EventManager) Notify(event CustomerEvent, customer *Customer) {
	for _, listener := range e.listeners {
		if listener.IsSubscribed(event) {
			listener.Notify(customer)
		}
	}
}
