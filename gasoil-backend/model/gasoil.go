package model

type GasResponse struct {
	Fuels		[]GasPrice	`json:"fuels"`
	UpdatedAt	string		`json:"updated_at"`
	URL			string		`json:"url"`
}

type GasPrice struct {
	Type	string 	`json:"type"`
	Price	float64 	`json:"price"`
}