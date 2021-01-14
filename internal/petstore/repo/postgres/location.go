package postgres

import (
	"fmt"
	client "petstore/internal/database/postgres"
	"petstore/pkg/models"
)

//Location use from models
type Location models.Location
type LocationService models.LocationService

//GetAllLocation get all Locations
func (l Location) GetAllLocation() ([]*models.Location, error) {
	var locs []*models.Location
	rows, err := client.DbClient.Query("SELECT * from public.Location")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&l.ID, &l.LocationName)
		cm := models.Location(l)
		locs = append(locs, &cm)
		if err != nil {
			return nil, err
		}

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return locs, nil
}

//CreateLocation create a Location
func (l *Location) CreateLocation() error {
	fmt.Println(client.DbClient)
	stmt, err := client.DbClient.Prepare("INSERT INTO public.location (location_name) VALUES($1,$2);")
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
