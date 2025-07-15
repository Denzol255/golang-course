package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var operations = map[string]func([]int) float64{
	"SUM": calculateSum,
	"AVG": calculateAvg,
	"MED": calculateMed,
}

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

	if operationFunc, ok := operations[operation]; ok {
		fmt.Printf("Результат: %.2f\n", operationFunc(numbers))
		return
	}
	fmt.Println("Некорректная операция")
}

func calculateSum(numbers []int) float64 {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return float64(sum)
}

func calculateAvg(numbers []int) float64 {
	return calculateSum(numbers) / float64(len(numbers))
}

func calculateMed(numbers []int) float64 {
	sort.Ints(numbers)
	if len(numbers)%2 == 0 {
		return (float64(numbers[len(numbers)/2]) + float64(numbers[len(numbers)/2-1])) / 2.0
	}
	return float64(numbers[(len(numbers)-1)/2])
}
