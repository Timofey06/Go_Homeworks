package dto

type TransferRequest struct {
	SenderId uint64 `json:"senderId"`
	TargetId uint64 `json:"targetId"`
	Value    int64  `json:"value"`
}
