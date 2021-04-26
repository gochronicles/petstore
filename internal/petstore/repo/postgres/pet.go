package postgres

import (
	"database/sql"
	"errors"
	client "petstore/internal/database/postgres"
	"petstore/pkg/models"
)

//Pet use from models
type Pet models.Pet

//PetService use from models
type PetService models.PetService

//GetPet get one Pet
func (p Pet) GetPet(id int) (*models.Pet, error) {
	row := client.DbClient.QueryRow("select pet.id,pet.pet_name,pet.age,pet.description,pet.category_id,category.category_name,pet.breed_id,breed.breed_name,pet.location_id,location.location_name from pet inner join category on pet.category_id = category.id inner join breed on pet.breed_id=breed.id inner join location on pet.location_id = location.id where pet.id=$1;", id)
	switch err := row.Scan(&p.ID, &p.Name, &p.Age, &p.Description, &p.PetCategory.ID, &p.PetCategory.CategoryName, &p.PetBreed.ID, &p.PetBreed.BreedName, &p.PetLocation.ID, &p.PetLocation.LocationName); err {
	case sql.ErrNoRows:
		return nil, errors.New("No matching records")
	case nil:
		category := models.Pet(p)
		return &category, nil
	default:
		return nil, err
	}
}

//GetPetByCategory based on category, get all Pets
func (p Pet) GetPetByCategory(categoryID int) ([]*models.Pet, error) {
	var pets []*models.Pet
	rows, err := client.DbClient.Query("select pet.id,pet.pet_name,pet.age,pet.description,pet.category_id,category.category_name,pet.breed_id,breed.breed_name,pet.location_id,location.location_name from pet inner join category on pet.category_id = category.id inner join breed on pet.breed_id=breed.id inner join location on pet.location_id = location.id where category.id=$1;", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Age, &p.Description, &p.PetCategory.ID, &p.PetCategory.CategoryName, &p.PetBreed.ID, &p.PetBreed.BreedName, &p.PetLocation.ID, &p.PetLocation.LocationName)
		pm := models.Pet(p)
		pets = append(pets, &pm)
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, err
	}

	return pets, nil
}

//CreatePet create a Pet
func (p *Pet) CreatePet() error {
	stmt, err := client.DbClient.Prepare("INSERT INTO pet (name,age,image_url,description,breed_id,category_id,location_id) VALUES($1,$2,$3,$4,$5,$6,$7);")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Age, p.ImageURL, p.Description, p.PetBreed.ID, p.PetCategory.ID, p.PetLocation.ID)
	if err != nil {
		return err
	}
	return nil
}

//DeletePet delete a Pet
func (p *Pet) DeletePet(id int) error {
	stmt, err := client.DbClient.Prepare("DELETE from pet where id=$1")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
