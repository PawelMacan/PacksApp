package domain

type InputRequest struct {
	Amount int `json:"amount"`
}

type CalculationResult struct {
	Packs           map[int]int `json:"packs"`
	TotalItems      int         `json:"total_items"`
	RequestedAmount int         `json:"requested_amount"`
	Overage         int         `json:"overage"`
	TotalPacks      int         `json:"total_packs"`
}
