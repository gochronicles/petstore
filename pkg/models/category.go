package models

//Category model to store category
type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryname"`
}

//CategoryService interface for Category model
type CategoryService interface {
	GetCategory(id int) (*Category, error)
	GetAllCategory() ([]*Category, error)
	CreateCategory() error
}
