package main

import "fmt"

func main() {
	const USDtoEUR float64 = 0.88
	const USDtoRUB float64 = 79.0
	const EURtoRUB float64 = USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
}
