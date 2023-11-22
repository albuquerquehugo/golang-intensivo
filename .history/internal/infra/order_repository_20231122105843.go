package infra

import (
	"database/sql"

	"github.com/albuquerquehugo/golang-intensivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec(`INSERT INTO order VAL)
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {

}