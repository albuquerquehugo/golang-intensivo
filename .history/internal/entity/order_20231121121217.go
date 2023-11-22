package entity

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price)

func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}