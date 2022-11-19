package model

type user struct {
	ID      int64   `json:"id"`
	Balance float64 `json:"balance"`
}

func NewUser() *user {
	return &user{
		ID:      0,
		Balance: 0.0,
	}
}

func (u *user) SetBalance(amount float64) {
	u.Balance = amount
}
