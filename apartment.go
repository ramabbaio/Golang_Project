package apartment

import "fmt"

type Apartment struct {
	address string
	square  float64
	floor   int
	price   int
}

func NewApartment(
	address string,
	square float64,
	floor int,
	price int,
) Apartment {
	if address == "" {
		fmt.Println("Адрес не может быть пустым, ошибка валидации")
		return Apartment{}
	}

	if square < 10 {
		fmt.Println("Площадь не может быть меньше 10, ошибка")
		return Apartment{}
	}

	if floor < 0 || floor > 50 {
		fmt.Println("Недопустимый этаж, ошибка")
		return Apartment{}
	}

	if price < 0 {
		fmt.Println("Цена не может быть меньше 0, ошибка")
		return Apartment{}
	}

	return Apartment{
		address: address,
		square:  square,
		floor:   floor,
		price:   price,
	}
}

func (a *Apartment) CostUp(cost int) {
	if cost < 0 {
		fmt.Println("Ошибка повышения цены на отрицательный")
		return
	}
	a.price += cost
}
