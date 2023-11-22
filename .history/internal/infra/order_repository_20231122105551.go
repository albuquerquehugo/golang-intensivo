package infra

import "database/sql"

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *Order) error {

}

GetTotal() (int64, error) {

}