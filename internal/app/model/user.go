package model

type User struct {
	ID      int64   `db:"id"`
	Balance float64 `db:"balance"`
}
