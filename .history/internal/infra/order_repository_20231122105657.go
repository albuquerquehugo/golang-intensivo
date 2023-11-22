package infra

import "database/sql"

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *entity.Order) error {

}

func (r *OrderRepository) GetTotalTransactions() (int, error) {

}