package postgres

import (
	"database/sql"
	"errors"
	client "petstore/internal/database/postgres"
	"petstore/pkg/models"
)

//Category : import category object from models
type Category models.Category

//CategoryService : import category interface from smodels
type CategoryService models.CategoryService

//CreateCategory : create a new Category in DB
func (c *Category) CreateCategory() (int, error) {
	//Create and execute the query
	row := client.DbClient.QueryRow("INSERT INTO category (category_name) values ($1) RETURNING id;", c.CategoryName)
	// save the data in an object
	switch err := row.Scan(&c.ID); err {
	// success, no errors
	case nil:
		return c.ID, nil
	// In case of no output rows
	case sql.ErrNoRows:
		return c.ID, errors.New("Insert Failed")
	// other errors occured
	default:
		return c.ID, err
	}
}

//GetCategory : get one Category based on an id in the DB
func (c Category) GetCategory(id int) (*models.Category, error) {
	// create and execute query
	row := client.DbClient.QueryRow("SELECT * from public.category WHERE id=$1", id)
	switch err := row.Scan(&c.ID, &c.CategoryName); err {
	// no matching records found
	case sql.ErrNoRows:
		return nil, errors.New("No records found")
	// success case
	case nil:
		// type cast explicitly to object of models.Category as
		// interface defined returns model.Category
		category := models.Category(c)
		return &category, nil
	//other errors
	default:
		return nil, err
	}
}

//GetAllCategory : Get all Category from DB
func (c Category) GetAllCategory() ([]*models.Category, error) {
	var categories []*models.Category // this will hold final result
	rows, err := client.DbClient.Query("SELECT * from public.category")
	defer rows.Close()
	// iterate through rows
	for rows.Next() {
		// bind to model
		err = rows.Scan(&c.ID, &c.CategoryName)
		// type cast to models.Category
		cm := models.Category(c)
		//add to result
		categories = append(categories, &cm)
		if err != nil {
			return nil, err
		}

	}
	// handle any errors that are encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
