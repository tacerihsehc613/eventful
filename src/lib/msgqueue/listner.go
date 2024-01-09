package msgqueue

type EventListener interface {
	Listen(events ...string) (<-chan Event, <-chan error, error)
	//Listen(stopChan <-chan struct{}, events ...string) (<-chan Event, <-chan error, error)
}
