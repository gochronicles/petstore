package postgres

import (
	"petstore/pkg/models"
)

//Category use from models
type Category models.Category

//CategoryService w
type CategoryService models.CategoryService

//GetCategory get one Category
func (c *Category) GetCategory(id int) (*models.Category, error) {
	return nil, nil
}

//GetAllCategory get all Categorys
func (c *Category) GetAllCategory() ([]*models.Category, error) {
	return nil, nil
}

//CreateCategory create a Category
func (c *Category) CreateCategory() error {
	stmt, err := dbClient.Prepare("INSERT INTO public.category (category_name) VALUES(?);")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()

	_, err = stmt.Exec(c.CategoryName)

	if err != nil {
		return err
	}
	return nil
}

//DeleteCategory delete a Category
func (c *Category) DeleteCategory(id int) error {
	return nil
}
