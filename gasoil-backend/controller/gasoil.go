package controller

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/HayataTakagi/GasoilCalculator/model"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type GasOil struct {}

func NewGaOil() *GasOil {
	return &GasOil{}
}

func (g *GasOil) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	gasRes := model.GasResponse{}
	url := "https://gogo.gs/ranking/average/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	priceSelections := doc.Find("div.pref-price-panel > div.average")
	updated_at := doc.Find("div.pref-price-panel > span.help").Text()
	gasRes.UpdatedAt = parseDate(updated_at)
	gasRes.URL = url
	priceSelections.Each(func(index int, s *goquery.Selection) {
		oilType := s.Find("label").Text()
		stringPrice := s.Find("div.price").Text()
		gasPrice := model.GasPrice{}
		gasPrice.Type = replaceType(oilType)
		gasPrice.Price = parsePrice(stringPrice)
		setFuel(gasPrice, &gasRes.Fuels)
		if err != nil {
			panic(err)
		}
	})
	fmt.Fprintf(os.Stdout ,"Res: %v\n", gasRes)
	return http.StatusOK, gasRes, nil
}

func parseDate(date string) string {
	parsedDate := strings.Split(date, "（")
	removedSpaceDate := strings.TrimSpace(parsedDate[0])
	r := strings.NewReplacer("年", "/", "月", "/", "日", "",)
	replacedDate := r.Replace(removedSpaceDate)
	return replacedDate
}

func parsePrice(stringPrice string) float64  {
	replacedPrice := strings.Replace(stringPrice, ",", "", -1)
	exchangedPrice, err := strconv.ParseFloat(replacedPrice, 64)
	if err != nil {
		panic(err)
	}
	return exchangedPrice
}

func replaceType(JapType string) string {
	r := strings.NewReplacer(
		"レギュラー", "regular",
		"ハイオク", "high_octane",
		"軽油", "light",
		"灯油", "kerosene",
		)
	return r.Replace(JapType)
}

func setFuel(price model.GasPrice, fuels *model.Fuels) {
	switch price.Type {
	case "regular":
		fuels.Regular = price
	case "high_octane":
		fuels.HighOctane = price
	case "light":
		fuels.Light = price
	case "kerosene":
		fuels.Kerosene = price
	}
}