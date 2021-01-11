package service

import (
	"fmt"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/pkg/models"
)

var err error
var bs db.BreedService

//CreateBreed create a breed
func CreateBreed(b *db.Breed) error {
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
