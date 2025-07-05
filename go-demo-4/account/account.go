package account

import (
	"errors"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) GetInfo() {
	blue := color.New(color.FgBlue).PrintfFunc()

	blue("Логин: %s\n", acc.Login)
	blue("Пароль: %s\n", acc.Password)
	blue("URL: %s\n", acc.Url)
	blue("Дата создания: %s\n", acc.CreatedAt.Format(time.DateTime))
	blue("Дата обновления: %s\n", acc.UpdatedAt.Format(time.DateTime))
}

func (acc *Account) generatePassword(length int) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?"
	var password string
	for range length {
		password += string(charset[rand.Intn(len(charset))])
	}
	acc.Password = password
}

func NewAccount(login, password, urlValue string) (*Account, error) {
	// Валидация логина
	if login == "" {
		return nil, errors.New("логин не может быть пустым")
	}
	// Валидация url
	_, err := url.ParseRequestURI(urlValue)
	if err != nil {
		return nil, err
	}

	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlValue,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if newAcc.Password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}
