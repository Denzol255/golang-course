package main

import (
	"app/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Green(`
		Добро пожаловать в программу для хранения паролей!
		Выберите действие:
	`)
	vault := account.NewVault()
Menu:
	for {
		switch getUserChoice() {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAccount(vault *account.Vault) {
	login := getPromptData("Введите логин: ")
	password := getPromptData("Введите пароль: ")
	url := getPromptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		color.Red(err.Error())
		return
	}

	vault.AddAccount(*myAccount)

}

func findAccount(vault *account.Vault) {
	url := getPromptData("Введите url: ")
	if len(url) == 0 {
		color.Red("URL не может быть пустым")
	} else {
		result := vault.FindAccountsByUrl(url)
		fmt.Println()
		if len(result) == 0 {
			color.Red("Не найдено ни одного аккаунта")
		} else {
			if len(result) > 1 {
				color.Yellow("Найдено %d аккаунтов", len(result))
			} else {
				color.Yellow("Найден 1 аккаунт")
			}
			for index, value := range result {
				fmt.Println()
				color.Green("%d. Данные аккаунта", index+1)
				value.GetInfo()
				fmt.Println()
			}
		}
		fmt.Println()
	}
}

func deleteAccount(vault *account.Vault) {
	url := getPromptData("Введите url: ")
	if len(url) == 0 {
		color.Red("URL не может быть пустым")
	} else {
		found := vault.DeleteAccountByUrl(url)
		if found {
			color.Green("Аккаунт успешно удален")
		} else {
			color.Red("Не найдено ни одного аккаунта")
		}
	}
}

func getPromptData(prompt string) string {
	var value string
	color.Green(prompt)
	fmt.Scanln(&value)
	return value
}

func getUserChoice() int {
	var choice int
	color.Green("Выберите действие: ")
	color.Green("1. Создать аккаунт")
	color.Blue("2. Найти аккаунт")
	color.Yellow("3. Удалить аккаунт")
	color.White("4. Выход")
	fmt.Scanln(&choice)
	return choice
}

func reverse(pointer *[5]int) {
	for index, element := range *pointer {
		(*pointer)[len(*pointer)-index-1] = element
		fmt.Println(index, element)
		fmt.Println(*pointer)
	}
}

func reverseWithoutCopy(a *[5]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
