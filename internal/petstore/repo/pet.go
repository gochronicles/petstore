package repo

import (
	"petstore/pkg/models"
)

//Pet use from models
type Pet models.Pet
type PetService models.PetService

//GetPet get one Pet
func (p *Pet) GetPet(id int) (*models.Pet, error) {
	return nil, nil
}

//GetPetByCategory based on category, get all Pets
func (p *Pet) GetPetByCategory(categoryID int) ([]*models.Pet, error) {
	return nil, nil
}

//CreatePet create a Pet
func (p *Pet) CreatePet() error {
	return nil
}

//DeletePet delete a Pet
func (p *Pet) DeletePet(id int) error {
	return nil
}
