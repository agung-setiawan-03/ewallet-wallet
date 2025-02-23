package models

type TransactionRequest struct {
	Reference string  `json:"reference"`
	Amount    float64 `json:"amount"`
}

type TransactionResponse struct {
	Reference string  `json:"reference"`
	Balance    float64 `json:"balance"`
}