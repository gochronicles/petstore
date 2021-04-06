package models

//Pet storage model for pet
type Pet struct {
	ID          int      `json:"id" form:"id"`
	Name        string   `json:"name" form:"name"`
	Age         float64  `json:"age" form:"age"`
	Description string   `json:"description" form:"description"`
	ImageURL    string   `json:"imageUrl" form:"imageUrl"`
	PetLocation Location `json:"location" form:"location"`
	PetCategory Category `json:"category" form:"category"`
	PetBreed    Breed    `json:"breed" form:"breed"`
}

//PetService interface for Pet model
type PetService interface {
	CreatePet() error
	GetPetByCategory(categoryID int) ([]*Pet, error)
	GetPet(id int) (*Pet, error)
	DeletePet(id int) error
	UpdateImageURL() error
}
