package entity

type Order struct {
	ID string
	Price float64
	Tax
	FinalPrice float
}

