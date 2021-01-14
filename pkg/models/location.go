package models

//Location model for location
type Location struct {
	ID           int    `json:"id"`
	LocationName string `json:"locationName"`
}

//LocationService interface for location model
type LocationService interface {
	CreateLocation() error
	GetAllLocation() ([]*Location, error)
}
