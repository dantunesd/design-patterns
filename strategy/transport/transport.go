package transport

type Transport interface {
	Transport(person string)
	GetType() transportType
}

type transportType string
