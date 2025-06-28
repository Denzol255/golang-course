package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("__Welcome to IMT Calculator__")
	fmt.Println(reverseString("Привет"))
	fmt.Println(substr("hello", 1, 4))
	fmt.Println(toUpperCase("hello, dude"))
	replaceString()
	weight, height := getUserData()
	IMT := calculateIMT(weight, height)
	fmt.Printf("Your IMT is: %.0f\n", IMT)
}

func calculateIMT(weight float64, height float64) float64 {
	const IMTPower = 2
	return weight / math.Pow(height/100, IMTPower)
}

func getUserData() (float64, float64) {
	var userWeight float64
	var userHeight float64
	fmt.Println("What is your weight in kilograms?")
	fmt.Scan(&userWeight)
	fmt.Println("What is your height in centimeters?")
	fmt.Scan(&userHeight)
	return userWeight, userHeight
}

func reverseString(s string) string {
	// Преобразовать строку в срез рун
	runes := []rune(s)
	// Перевернуть срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Преобразовать срез рун обратно в строку
	return string(runes)
}

func substr(str string, start, end int) string {
	return str[start:end]
}

func toUpperCase(input string) string {
	var result []rune
	for _, r := range input {
		result = append(result, unicode.ToUpper(r))
	}
	return string(result)
}

func replaceString() {
	str := "Привет, мир"
	old := "мир"
	new := "Go"
	// Меняем все вхождения подстроки old на new в строке str
	result := strings.Replace(str, old, new, -1)
	fmt.Println(result)
}
