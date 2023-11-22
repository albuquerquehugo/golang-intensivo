package main

type Car struct {
	Model string
	Color string
}

func soma(x,y int) int {

}

func main() {
	car := Car {
		Model: "Ferrari",
		Color: "Red",
	}
	car.Model = "Fiat"
	println(car.Model)
}