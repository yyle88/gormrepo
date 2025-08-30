package models

type User struct {
	ID   uint
	Name string
}

func (*User) TableName() string {
	return "users"
}

type Order struct {
	ID     uint
	UserID uint
	Amount float64
}

func (*Order) TableName() string {
	return "orders"
}
