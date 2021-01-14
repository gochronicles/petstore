package service

import (
	db "petstore/internal/petstore/repo/postgres"
	"petstore/pkg/models"
)

var cs db.CategoryService
var category db.Category

//CreateCategory create a breed
func CreateCategory(c *db.Category) error {
	cs = c
	err = cs.CreateCategory()
	if err != nil {
		return err
	}
	return nil
}

//GetAllCategory get breed based on category
func GetAllCategory() ([]*models.Category, error) {
	cs = &category
	categories, err := cs.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

//GetCategory get
func GetCategory(id int) (*models.Category, error) {
	cs = &category
	c, err := cs.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
