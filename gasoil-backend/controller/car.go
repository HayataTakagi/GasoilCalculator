package controller

import (
	"github.com/jmoiron/sqlx"
	"github.com/HayataTakagi/treasure-app/model"
	"github.com/HayataTakagi/treasure-app/repository"
	"net/http"
)

type Car struct {
	db *sqlx.DB
}

func NewCar(db *sqlx.DB) *Car {
	return &Car{db: db}
}

func (c *Car) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	cars, err := repository.AllCar(c.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	carRes := model.CarResponse{}
	url := "https://e-nenpi.com/"
	carRes.URL = url
	carRes.Cars = cars
	carRes.UpdatedAt = "2019/08/19"
	carRes.Count = int64(len(cars))
	return http.StatusOK, carRes, nil
}