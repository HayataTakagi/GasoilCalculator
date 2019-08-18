package controller

import (
	"fmt"
	"github.com/voyagegroup/treasure-app/model"
	"net/http"
	"os"
)

type Car struct {}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	carRes := model.CarResponse{}
	url := "https://e-nenpi.com/"
	cars := []model.Car{{Name: "トヨタ ヴィッツ（ハイブリッド）", FuelEconomy: 26.70},{Name: "トヨタ プリウス", FuelEconomy: 24.65},}
	carRes.URL = url
	carRes.Cars = cars
	carRes.Count = int64(len(cars))
	fmt.Fprintf(os.Stdout ,"Res: %v\n", carRes)
	return http.StatusOK, carRes, nil
}