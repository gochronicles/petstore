package service

import (
	db "petstore/internal/petstore/repo/postgres"
	"petstore/pkg/models"
)

var ls db.LocationService
var loc db.Location

//CreateLocation create a breed
func CreateLocation(l *db.Location) error {
	ls = l
	err = ls.CreateLocation()
	if err != nil {
		return err
	}
	return nil
}

//GetAllLocation get breed based on Location
func GetAllLocation() ([]*models.Location, error) {
	ls = &loc
	locs, err := ls.GetAllLocation()
	if err != nil {
		return nil, err
	}
	return locs, nil
}
