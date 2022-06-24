package customer

type Listener interface {
	Notify(customer *Customer)
	IsSubscribed(event CustomerEvent) bool
}

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
