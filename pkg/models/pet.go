package models

//Pet storage model for pet
type Pet struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Age         float64  `json:"age"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	PetLocation Location `json:"location"`
	PetCategory Category `json:"category"`
	PetBreed    Breed    `json:"breed"`
}

//PetService interface for Pet model
type PetService interface {
	CreatePet() error
	GetPetByCategory(categoryID int) ([]*Pet, error)
	GetPet(id int) (*Pet, error)
	DeletePet(id int) error
}
