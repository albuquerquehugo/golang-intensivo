package main

import (
	"database/sql"

	"github.com/albuquerquehugo/golang-intensivo/internal/entity"
	"github.com/albuquerquehugo/golang-intensivo/internal/usecase"
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

	db, err := sql.Open("sqlite3", "file")
	usecase := usecase.CalculateFinalPrice {
		OrderRepository: &database.OrderRepository{
			Db: db,
		},
	}
	input := entity.OrderInput{
		ID: "123",
		Price: 10.0,
		Tax: 1.0,
	}
}