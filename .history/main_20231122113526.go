package main

import (
	"database/sql"

	"github.com/albuquerquehugo/golang-intensivo/internal/infra/database"
	"github.com/albuquerquehugo/golang-intensivo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

// método
func (c Car) Start() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

// função comum
func soma(x, y int) int {
	return x + y
}

func main() {
	// car := Car {
	// 	Model: "Ferrari",
	// 	Color: "Red",
	// }
	// car.Model = "Fiat"
	// println(car.Model)
	// car.Start()
	// car.ChangeColor("Blue")
	// println(car.Color)

	// order, err := entity.NewOrder("1", -10, 1)
	// if err != nil {
	// 	println("Error creating order: ", err.Error())
	// } else {
	// 	println("Order #" + order.ID + " created successfully")
	// }

	// Criar tabela no banco de dados
	// > sqlite3 db.sqlite3
	// CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(&input)
	if err != nil {
		panic(err)
	}
	println(output)
}
