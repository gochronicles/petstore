package service

import (
	db "petstore/internal/petstore/repo/postgres"
	"petstore/pkg/models"
)

var cs db.CategoryService
var category db.Category

//CreateCategory : Create a category. calls respective DB or repo layer
func CreateCategory(c *db.Category) (int, error) {
	cs = c
	// call repo function
	id, err := cs.CreateCategory()
	if err != nil {
		return id, err
	}
	return id, nil
}

//GetAllCategory : Get all categories. calls respective DB or repo layer
func GetAllCategory() ([]*models.Category, error) {
	cs = &category
	// call repo function
	categories, err := cs.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

//GetCategory : Get one category based on id. calls respective DB or repo layer
func GetCategory(id int) (*models.Category, error) {
	cs = &category
	// call repo function
	c, err := cs.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
