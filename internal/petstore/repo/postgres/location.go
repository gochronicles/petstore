package postgres

import (
	"fmt"
	"petstore/pkg/models"
)

//Location use from models
type Location models.Location

//GetAllLocation get all Locations
func (l *Location) GetAllLocation() ([]*Location, error) {
	return nil, nil
}

//CreateLocation create a Location
func (l *Location) CreateLocation() error {
	fmt.Println(dbClient)
	stmt, err := dbClient.Prepare("INSERT INTO public.location (location_name) VALUES($1,$2);")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()
	_, err = stmt.Exec(l.LocationName)
	if err != nil {
		return err
	}
	return nil
}

//DeleteLocation delete a Location
func (l *Location) DeleteLocation(id int) error {
	return nil
}
