package domain

type Order struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type OrderList struct {
	Orders []Order `json:"orders"`
}