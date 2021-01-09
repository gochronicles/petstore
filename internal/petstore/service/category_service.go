package service

import (
	"petstore/internal/petstore/repo"
)

var(
	CategoryService = categoryservice{}
)

type categoryservice struct{}

func Save() error{
	category := repo.Category{
		Category_name: "Cat",
	}
	if err := category.Save(); err != nil {
		return err
	}
	return nil
}