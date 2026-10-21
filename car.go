package main

import "fmt"

type Car struct {
	mark   string
	model  string
	power  int
	speed  int
	razgon float64
	stage  int
}

func NewCar(
	mark string,
	model string,
	power int,
	speed int,
	razgon float64,
	stage int,
) Car {

	if mark == "" {
		fmt.Println("Пустая марка валидацию не прошла")
		return Car{}
	}

	if model == "" {
		fmt.Println("Пустая модель, валидацию не прошла")
		return Car{}
	}

	if power == 0 {
		fmt.Println("Сила меньше 0 не прошла валидацию")
		return Car{}
	}

	if speed > 300 || speed < 0 {
		fmt.Println("Невозможная скорость")
		return Car{}
	}

	if razgon < 1 || razgon > 20 {
		fmt.Println("Разгон не прошел валидацию")
		return Car{}
	}
	if stage < 0 || stage > 3 {
		return Car{}
	}

	return Car{
		mark:   mark,
		model:  model,
		power:  power,
		speed:  speed,
		razgon: razgon,
		stage:  stage,
	}
}

func (c *Car) PowerUp(power int) {
	if power < 0 || c.power+power > 1000 {
		fmt.Println("Ошибка не правильный ввод силы")
		return
	}
	c.power += power
}

func (c *Car) StageCar(level int) {
	if level < 0 || level > 3 {
		fmt.Println("Возможен стейдж от 0 до 3")
		return
	}

	switch level {
	case 1:
		c.power += 50
		c.speed += 5
		c.razgon -= 0.5
	case 2:
		c.power += 100
		c.speed += 10
		c.razgon -= 1
	case 3:
		c.power += 150
		c.speed += 20
		c.razgon -= 1.5
	}

	c.stage = level
}

func main() {
	camry := NewCar(
		"Лексус",
		"LX500",
		250,
		150,
		6.5,
		0,
	)
	fmt.Println("До вывода", camry)
	// camry.PowerUp(100)
	camry.StageCar(2)
	fmt.Println("После функций", camry)
}
