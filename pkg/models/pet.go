package models

//Pet storage model for pet
type Pet struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Age         float64 `json:"age"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
	LocationID  int     `json:"locationId"`
	CategoryID  int     `json:"categoryId"`
	BreedID     int     `json:"breedId"`
}

//PetService interface for Pet model
type PetService interface {
	CreatePet() error
	DeletePet(id int) error
	GetPetByCategory(categoryID int) ([]*Pet, error)
	GetPet(id int) (*Pet, error)
}
