package models

//Pet storage model for pet
type Pet struct {
	ID          int
	Name        string
	Age         float64
	Description string
	ImageURL    string
	LocationID  int
	CategoryID  int
	BreedID     int
}

//PetService interface for Pet model
type PetService interface {
	CreatePet() error
	DeletePet(id int) error
	GetPetByCategory(categoryID int) ([]*Pet, error)
	GetPet(id int) (*Pet, error)
}
