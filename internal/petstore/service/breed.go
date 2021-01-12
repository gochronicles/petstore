package service

import (
	"fmt"
	"petstore/internal/petstore/repo"
	"petstore/pkg/models"
)

var err error
var bs repo.BreedService

//CreateBreed create a breed
func CreateBreed(b *repo.Breed) error {
	fmt.Println("In service")
	bs = b
	err = bs.CreateBreed()
	if err != nil {
		return err
	}
	return nil
}

//DeleteBreed deletes a breed
func DeleteBreed(id int) error {
	return nil
}

//GetBreedByCategory get breed based on category
func GetBreedByCategory(categoryID int) ([]*models.Breed, error) {
	return nil, nil
}
