package models

//Breed : Model for storing Breed of pet
type Breed struct {
	ID         int    `json:"breedId"`
	BreedName  string `json:"breedName"`
	CategoryID int    `json:"categoryId"`
}

//BreedService interface for Breed struct
type BreedService interface {
	// create a new breed
	CreateBreed() error
	// get breed based on category identifiers
	GetBreedByCategory(categoryID int) ([]*Breed, error)
}
