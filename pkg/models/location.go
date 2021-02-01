package models

//Location : Model for storing Location
type Location struct {
	ID           int    `json:"id"`
	LocationName string `json:"locationName"`
}

//LocationService : Interface for location model
type LocationService interface {
	// add a new location
	CreateLocation() error
	// get all location
	GetAllLocation() ([]*Location, error)
}
