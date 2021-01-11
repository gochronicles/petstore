package postgres

import (
	"fmt"
	"petstore/pkg/models"
)

//Breed use from models
type Breed models.Breed

//BreedService use from models
type BreedService models.BreedService

//CreateBreed create a breed
func (b *Breed) CreateBreed() error {
	fmt.Println("In Repo")
	//Prepare statement for insert query
	return nil
}

//DeleteBreed delete a breed
func (b *Breed) DeleteBreed(id int) error {
	return nil
}

//GetBreedByCategory get breed based on category
func (b *Breed) GetBreedByCategory(categoryID int) ([]*models.Breed, error) {
	return nil, nil
}
