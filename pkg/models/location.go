package models

//Location model for location
type Location struct {
	ID           int
	LocationName string
}

//LocationService interface for location model
type LocationService interface {
	CreateLocation()
	DeleteLocation()
	GetAllLocation()
}
