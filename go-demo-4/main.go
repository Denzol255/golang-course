package main

import (
	"app/password/account"
	"app/password/files"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	color.Green(`
		Добро пожаловать в программу для хранения паролей!
		Выберите действие:
	`)
	vaultWithDB := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		variant := getPromptData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по login",
			"4. Удалить аккаунт",
			"5. Выход",
			"Введите номер действия",
		)
		callback := menu[variant]
		if callback == nil {
			break Menu
		}
		callback(vaultWithDB)
	}
}

func createAccount(vault *account.VaultWithDB) {
	login := getPromptData("Введите логин")
	password := getPromptData("Введите пароль")
	url := getPromptData("Введите url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		color.Red(err.Error())
		return
	}

	vault.AddAccount(*myAccount)

}

func findAccountByUrl(vault *account.VaultWithDB) {
	url := getPromptData("Введите URL")
	if len(url) == 0 {
		color.Red("URL не может быть пустым")
		return
	}

	result := vault.FindAccounts(func(account account.Account) bool {
		return strings.Contains(account.Url, url)
	})
	fmt.Println()
	displayFindingResult(&result)
	fmt.Println()
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := getPromptData("Введите login")
	if len(login) == 0 {
		color.Red("login не может быть пустым")
		return
	}

	result := vault.FindAccounts(func(account account.Account) bool {
		return strings.Contains(account.Login, login)
	})
	fmt.Println()
	displayFindingResult(&result)
	fmt.Println()
}

func displayFindingResult(result *[]account.Account) {
	if len(*result) == 0 {
		color.Red("Не найдено ни одного аккаунта")
		return
	}

	if len(*result) > 1 {
		color.Yellow("Найдено %d аккаунтов", len(*result))
	} else {
		color.Yellow("Найден 1 аккаунт")
	}

	for index, value := range *result {
		fmt.Println()
		color.Green("%d. Данные аккаунта", index+1)
		value.GetInfo()
		fmt.Println()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
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

func getPromptData(prompts ...any) string {
	var value string
	for index, prompt := range prompts {
		if index == len(prompts)-1 {
			color.Yellow("%v: ", prompt)
			break
		}
		color.Green("%v", prompt)
	}
	fmt.Scanln(&value)
	return value
}
