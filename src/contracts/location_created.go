package contracts

import "rabbit/lib/persistence"

type LocationCreatedEvent struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Country string             `json:"country"`
	Halls   []persistence.Hall `json:"halls"`
}

func (l *LocationCreatedEvent) EventName() string {
	return "location.created"
}
