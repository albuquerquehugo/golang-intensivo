package infra

import (
	"database/sql"

	"github.com/albuquerquehugo/golang-intensivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
		if err != nil {}
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {

}