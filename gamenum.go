package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1 // Создает число от 1 до 100 включительно
	attempt := 10 // 10 Попыток дается

	fmt.Println("--------Игра угадай число!--------")

	for {
		if attempt == 0 {
			fmt.Println("У вас закончились попытки, вы проиграли :(")
			break
		}

		fmt.Println("У вас есть: ", attempt, "Попыток")

		reader := bufio.NewReader(os.Stdin) // Делает обертку потока ввода 
		fmt.Print("Введите число: ")
		input, err := reader.ReadString('\n') // читает с буфера введенные сообщения до \n тоесть до enter

		if err != nil {
			log.Fatal(err) // отключает программу если error ошибка не равно пустоте тоесть если там есть ошибка выключает программу
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if num == target {
			fmt.Println("Вы угадали число:", target, "За", 10-attempt, "Попыток")
			break
		} else if num > target {
			fmt.Println("Число меньше!")
			fmt.Println()
			attempt--
		} else {
			fmt.Println("Число больше")
			fmt.Println()
			attempt--
		}

	}
}
