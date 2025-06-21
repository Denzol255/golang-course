package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/duke-git/lancet/v2/slice"
)

func main() {
	fmt.Println("Добро пожаловать в калькулятор!")
	var operation string
	var input string
	fmt.Print("Введите операцию: ")
	fmt.Scan(&operation)
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&input)

	stringNumbers := strings.Split(input, ",")
	numbers := make([]int, len(stringNumbers))
	for index, value := range stringNumbers {
		num, err := strconv.Atoi(value)
		if err == nil {
			numbers[index] = num
		} else {
			fmt.Println("Введено некорректное число: ", value)
			return
		}
	}
	var result float64
	switch operation {
	case "AVG":
		sum := slice.ReduceBy(numbers, 0, func(_, cur, acc int) int {
			return acc + cur
		})
		result = float64(sum) / float64(len(numbers))
	case "SUM":
		result = float64(slice.ReduceBy(numbers, 0, func(_, cur, acc int) int {
			return acc + cur
		}))
	case "MED":
		sort.Ints(numbers)
		if len(numbers)%2 == 0 {
			result = (float64(numbers[len(numbers)/2]) + float64(numbers[len(numbers)/2-1])) / 2.0
		} else {
			result = float64(numbers[(len(numbers)-1)/2])
		}
	default:
		fmt.Println("Некорректная операция")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
