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
	fmt.Println(dbClient)
	stmt, err := dbClient.Prepare("INSERT INTO public.breed (breed_name,category_id) VALUES($1,$2);")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()
	_, err = stmt.Exec(b.BreedName, b.CategoryID)
	if err != nil {
		return err
	}
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
