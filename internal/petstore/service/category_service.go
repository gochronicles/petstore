package service

import (
	"petstore/internal/petstore/repo"
)

var(
	CategoryService = categoryservice{}
)

type categoryservice struct{}

func (categoryservice categoryservice)Save(category *repo.Category) error{
	err := category.Save()
	if err != nil {
		return err
	}
	return nil
}