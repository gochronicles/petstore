package repo

import (
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
	//Prepare statement for insert query
	return nil
}

//DeleteLocation delete a Location
func (l *Location) DeleteLocation(id int) error {
	return nil
}
