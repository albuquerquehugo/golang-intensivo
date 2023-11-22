package main

type Car struct {
	Model string
	Color string
}

func soma(x,y int) int {
	return x+y
}

func main() {
	car := Car {
		Model: "Ferrari",
		Color: "Red",
	}
	car.Model = "Fiat"
	println(car.Model)
}