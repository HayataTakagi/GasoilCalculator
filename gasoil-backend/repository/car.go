package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/HayataTakagi/treasure-app/model"
)

func AllCar(db *sqlx.DB) ([]model.Car, error) {
	a := make([]model.Car, 0)
	if err := db.Select(&a, `SELECT id, name, fuel_economy FROM cars`); err != nil {
		return nil, err
	}
	return a, nil
}
