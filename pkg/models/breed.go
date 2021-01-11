package models

//Breed model for breed
type Breed struct {
	ID         int    `json:"breedId"`
	BreedName  string `json:"breedName"`
	CategoryID int    `json:"categoryId"`
}

//BreedService interface for breed
type BreedService interface {
	CreateBreed() error
	DeleteBreed(id int) error
	GetBreedByCategory(categoryID int) ([]*Breed, error)
}
