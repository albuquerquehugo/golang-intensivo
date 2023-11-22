package entity

type OrderRepositoryInterface interface {
	SaveOrder(order *Order) error
}