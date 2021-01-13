package postgres

import (
	"database/sql"
	"errors"
	client "petstore/internal/database/postgres"
	"petstore/pkg/models"
)

//Category use from models
type Category models.Category

//CategoryService w
type CategoryService models.CategoryService

//GetCategory get one Category
func (c Category) GetCategory(id int) (*models.Category, error) {
	row := client.DbClient.QueryRow("SELECT * from public.category WHERE id=$1", id)
	switch err := row.Scan(&c.ID, &c.CategoryName); err {
	case sql.ErrNoRows:
		return nil, errors.New("No matching records")
	case nil:
		category := models.Category(c)
		return &category, nil
	default:
		return nil, err
	}
}

//GetAllCategory get all Categorys
func (c Category) GetAllCategory() ([]*models.Category, error) {
	var categories []*models.Category
	rows, err := client.DbClient.Query("SELECT * from public.category")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.CategoryName)
		cm := models.Category(c)
		categories = append(categories, &cm)
		if err != nil {
			return nil, err
		}

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

//CreateCategory create a Category
func (c *Category) CreateCategory() error {
	stmt, err := client.DbClient.Prepare("INSERT INTO public.category (category_name) values ($1);")
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
