package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type currencyMap map[string]map[string]float64

const USDtoEUR float64 = 0.88
const USDtoRUB float64 = 79.0
const EURtoRUB float64 = USDtoRUB / USDtoEUR

var CURRENCY_MAP = currencyMap{
	"USD": {"EUR": USDtoEUR, "RUB": USDtoRUB},
	"EUR": {"RUB": EURtoRUB, "USD": 1 / USDtoEUR},
	"RUB": {"EUR": 1 / EURtoRUB, "USD": 1 / USDtoRUB},
}

func main() {
	fmt.Println("__Добро пожаловать в валютный калькулятор__")
	for {
		sum, currencyFrom, currencyTo, error := getUserData()

		if error != nil {
			fmt.Println(error)
			continue
		}

		result := convert(sum, currencyFrom, currencyTo, &CURRENCY_MAP)
		fmt.Printf("%.2f %s = %.2f %s\n", sum, currencyFrom, result, currencyTo)

		fmt.Println("Хотите продолжить работу с калькулятором? (y/n)")

		var answer string
		fmt.Scan(&answer)
		if answer == "n" {
			break
		}
	}

}

func getUserData() (float64, string, string, error) {
	var sum float64
	var currencyFrom string
	var currencyTo string

	fmt.Println("Введите исходную валюту (USD, EUR или RUB):")
	fmt.Scan(&currencyFrom)
	if currencyFrom != "USD" && currencyFrom != "EUR" && currencyFrom != "RUB" {
		return sum, currencyFrom, currencyTo, errors.New("Некорректная исходная валюта")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите сумму: ")
	input, error := reader.ReadString('\n')
	value, error := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if error != nil || value <= 0 {
		return sum, currencyFrom, currencyTo, errors.New("Сумма должна быть больше нуля")
	}
	sum = value

	if currencyFrom == "USD" {
		fmt.Println("Введите валюту для конвертации (EUR или RUB):")
	}
	if currencyFrom == "EUR" {
		fmt.Println("Введите валюту для конвертации (USD или RUB):")
	}
	if currencyFrom == "RUB" {
		fmt.Println("Введите валюту для конвертации (EUR или USD):")
	}
	fmt.Scan(&currencyTo)
	if currencyTo == currencyFrom {
		return sum, currencyFrom, currencyTo, errors.New("Введены одинаковые валюты")
	} else if currencyTo != "USD" && currencyTo != "EUR" && currencyTo != "RUB" {
		return sum, currencyFrom, currencyTo, errors.New("Некорректная валюта для конвертации")
	}

	return sum, currencyFrom, currencyTo, nil
}
func convert(sum float64, currencyFrom string, currencyTo string, mapPointer *currencyMap) float64 {
	return sum * (*mapPointer)[currencyFrom][currencyTo]
}
