package service

import (
	db "petstore/internal/petstore/repo/postgres"
	"petstore/pkg/models"
)

var ps db.PetService
var pet db.Pet

//CreatePet create a Pet
func CreatePet(p *db.Pet) error {
	ps = p
	err = ps.CreatePet()
	if err != nil {
		return err
	}
	return nil
}

//DeletePet deletes a Pet
func DeletePet(id int) error {
	ps = &pet
	err = ps.DeletePet(id)
	if err != nil {
		return err
	}
	return nil
}

//GetPet get one Pet
func GetPet(id int) (*models.Pet, error) {
	ps = &pet
	pet, err := ps.GetPet(id)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

//GetPetByCategory get Pet based on category
func GetPetByCategory(categoryID int) ([]*models.Pet, error) {
	ps = &pet
	pets, err := ps.GetPetByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return pets, nil
}
