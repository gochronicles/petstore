package repo

import (
	"errors"
	postgres2 "petstore/internal/postgres"
	"petstore/pkg/models"
)

type Category models.Category

func (category *Category) Save() error {

	if category == nil {
		return errors.New("Invlaid Category to save")
	}
	//Prepare statement for insert query
	stmt, err := postgres2.DbClient.Prepare("INSERT INTO public.category (category_name) VALUES(?);")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()

	_, err = stmt.Exec(category.Category_name)

	if err != nil {
		return err
	}

	return nil
}
