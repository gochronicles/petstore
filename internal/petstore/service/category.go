package service

import (
	db "petstore/internal/petstore/repo/postgres"
	"petstore/pkg/models"
)

var cs db.CategoryService

//CreateCategory create a breed
func CreateCategory(c *db.Category) error {
	cs = c
	err = cs.CreateCategory()
	if err != nil {
		return err
	}
	return nil
}

//DeleteCategory deletes a breed
func DeleteCategory(id int) error {
	return nil
}

//GetAllCategory get breed based on category
func GetAllCategory() ([]*models.Category, error) {
	return nil, nil
}

//GetCategory get
func GetCategory() (*models.Category, error) {
	return nil, nil
}
