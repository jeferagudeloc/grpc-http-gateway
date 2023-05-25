package domain

type Order struct {
	ID           string `json:"id"`
	OrderType    string `json:"orderType"`
	Store        string `json:"store"`
	Address      string `json:"address"`
	CreationDate string `json:"creationDate"`
	Status       string `json:"status"`
}
