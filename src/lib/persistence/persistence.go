package persistence

type DatabaseHandler interface {
	//AddBookingForUser([]byte, Booking) error
	AddBookingForUser(string, Booking) error

	//AddEvent(Event) ([]byte, error)
	AddEvent(Event) (string, string, error)
	AddEvent4Booking(Event) (string, string, error)
	//FindEvent([]byte) (Event, error)
	FindEvent(string) (Event, error)
	FindEvent4Booking(string) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)

	AddLocation(Location) (Location, error)
	FindLocation(string) (Location, error)
	FindAllLocations() ([]Location, error)
}
