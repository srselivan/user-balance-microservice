package model

type HolderStruct struct {
	UserID    int64   `json:"user_id"`
	OrderID   int64   `json:"order_id"`
	ServiceID int64   `json:"service_id"`
	Amount    float64 `json:"amount"`
}

func (h *HolderStruct) Valid() bool {
	return h.UserID >= 0 && h.OrderID >= 0 && h.ServiceID >= 0 && h.Amount >= 0
}
