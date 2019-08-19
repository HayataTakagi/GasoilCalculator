package model

type GasResponse struct {
	Fuels		Fuels		`json:"fuels"`
	UpdatedAt	string		`json:"updated_at"`
	URL			string		`json:"url"`
}

type Fuels struct {
	Regular 	GasPrice `json:"regular"`
	HighOctane 	GasPrice `json:"high_octane"`
	Light 		GasPrice `json:"light"`
	Kerosene 	GasPrice `json:"kerosene"`
}

type GasPrice struct {
	Type	string 		`json:"type"`
	Price	float64 	`json:"price"`
}