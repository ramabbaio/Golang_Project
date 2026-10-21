package main

import "fmt"

type Car struct {
	mark  string
	model string
	power int
	price int
}

func main() {
	camry := Car{
		mark:  "Тойота",
		model: "Camry",
		power: 250,
		price: 600000,
	}

	camry.mark = "Lexus"
	fmt.Println(camry)
}
