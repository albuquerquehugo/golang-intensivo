package infra

import (
	"database/sql"

	"github.com/albuquerquehugo/golang-intensivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err 
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {

}