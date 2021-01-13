package postgres

import (
	"fmt"
	client "petstore/internal/database/postgres"
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
	fmt.Println(client.DbClient)
	stmt, err := client.DbClient.Prepare("INSERT INTO public.pet (name,image_url,description,breed_id,category_id,location_id) VALUES(?);")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.ImageURL, p.Description, p.BreedID, p.CategoryID, p.LocationID)
	if err != nil {
		return err
	}
	return nil
}

//DeletePet delete a Pet
func (p *Pet) DeletePet(id int) error {
	return nil
}
