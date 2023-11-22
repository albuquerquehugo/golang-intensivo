package main

type Car struct {
	Model string
	Color string
}

// método
func (c Car) Start() {
	println(c.Model + "is started")
}

// função comum
func soma(x, y int) int {
	return x + y
}

func main() {
	car := Car {
		Model: "Ferrari",
		Color: "Red",
	}
	car.Model = "Fiat"
	println(car.Model)
	car.
}