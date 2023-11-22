package usecase

type OrderInput struct {
	ID string
	Price float64
	Tax float64
}

type OrderOutput struct {
	ID string
	Price float64
	Tax float64
	
}