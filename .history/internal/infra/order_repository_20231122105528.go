package infra

import "database/sql"

type OrderRepository struct {
	Db *sql.DB
}

Save(order *Order) error {

}

GetTotal