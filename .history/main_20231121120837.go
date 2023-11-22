package main

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
	a := 10
	b := a
	b = 20
	println(a)
	println(b)

	// car := Car {
	// 	Model: "Ferrari",
	// 	Color: "Red",
	// }
	// car.Model = "Fiat"
	// println(car.Model)
	// car.Start()
	// car.ChangeColor("Blue")
	// println(car.Color)
}