package persistence

type DatabaseHandler interface {
	AddBookingForUser([]byte, Booking) error

	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)

	AddLocation(Location) (Location, error)
	FindLocation(string) (Location, error)
	FindAllLocations() ([]Location, error)
}
