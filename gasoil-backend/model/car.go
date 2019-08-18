package model

type CarResponse struct {
	Count		int64		`json:"count"`
	Cars		[]Car		`json:"cars"`
	UpdatedAt	string		`json:"updated_at"`
	URL			string		`json:"url"`
}

type Car struct {
	ID			int64 	`db:"id" json:"id"`
	Name 		string 	`db:"name" json:"name"`
	FuelEconomy	float64 `db:"fuel_economy" json:"fuelEconomy"`
}