package service

import (
	"fmt"
	"petstore/internal/petstore/repo"
	"petstore/pkg/models"
)

var ps repo.PetService

//CreatePet create a Pet
func CreatePet(p *repo.Pet) error {
	fmt.Println("In service")
	ps = p
	err = ps.CreatePet()
	if err != nil {
		return err
	}
	return nil
}

//DeletePet deletes a Pet
func DeletePet(id int) error {
	return nil
}

//GetPet get one Pet
func GetPet(id int) (*models.Pet, error) {
	return nil, nil
}

//GetPetByCategory get Pet based on category
func GetPetByCategory(categoryID int) ([]*models.Pet, error) {
	return nil, nil
}
