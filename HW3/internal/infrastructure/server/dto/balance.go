package dto

type BalanceRequest struct {
	Id uint64 `json:"id"`
}

type BalanceResponse struct {
	Balance int64 `json:"balance"`
}
