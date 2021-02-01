package models

//Category : Model for storing category of the pet
type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryname"`
}

//CategoryService interface for Category model
type CategoryService interface {
	// create a new category
	CreateCategory() (int, error)
	// get a category based on id
	GetCategory(id int) (*Category, error)
	// get all categories
	GetAllCategory() ([]*Category, error)
}
