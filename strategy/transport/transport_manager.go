package transport

type transports map[transportType]Transport

type Manager struct {
	transports transports
}

func NewManager() *Manager {
	return &Manager{
		transports: make(transports),
	}
}

func (m *Manager) Add(transport Transport) {
	m.transports[transport.GetType()] = transport
}

func (m *Manager) Transport(person string, transportType transportType) {
	m.transports[transportType].Transport(person)
}
