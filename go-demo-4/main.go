package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *accountWithTimeStamp) getInfo() {
	fmt.Println("Информация о аккаунте:")
	fmt.Println("Логин:", acc.login)
	fmt.Println("Пароль:", acc.password)
	fmt.Println("URL:", acc.url)
	fmt.Println("Дата создания:", acc.createdAt.Format(time.RFC3339))
	fmt.Println("Дата обновления:", acc.updatedAt.Format(time.RFC3339))
}

func (acc *account) generatePassword(length int) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?"
	var password string
	for range length {
		password += string(charset[rand.Intn(len(charset))])
	}
	acc.password = password
}

func newAccountWithTimeStamp(login, password, urlValue string) (*accountWithTimeStamp, error) {
	// Валидация логина
	if login == "" {
		return nil, errors.New("логин не может быть пустым")
	}
	// Валидация url
	_, err := url.ParseRequestURI(urlValue)
	if err != nil {
		return nil, err
	}

	newAcc := &accountWithTimeStamp{
		account: account{login: login,
			password: password,
			url:      urlValue},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if newAcc.password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}

func main() {
	login := getPromptData("Введите логин: ")
	password := getPromptData("Введите пароль: ")
	url := getPromptData("Введите url: ")

	account, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	account.getInfo()
}

func getPromptData(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
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
