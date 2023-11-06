package msgqueue

// interface{} is an empty interface, which means that it can be satisfied by any type
type EventMapper interface {
	MapEvent(string, interface{}) (Event, error)
}

func NewEventMapper() EventMapper {
	return &StaticEventMapper{}
}
