package account

import (
	"app/password/files"
	"encoding/json"
	"slices"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vault *Vault) DeleteAccountByUrl(url string) bool {
	accountIdx := slices.IndexFunc(vault.Accounts, func(account Account) bool {
		return account.Url == url
	})
	if accountIdx != -1 {
		vault.Accounts = append(vault.Accounts[:accountIdx], vault.Accounts[accountIdx+1:]...)
		vault.save()
	}
	return accountIdx != -1
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	result := make([]Account, 0)
	for _, value := range vault.Accounts {
		if strings.Contains(value.Url, url) {
			result = append(result, value)
		}
	}

	return result
}

func NewVault() *Vault {
	file, err := files.ReadFromFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  make([]Account, 0),
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (vault *Vault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.save()
}

func (vault *Vault) save() {
	data, err := json.Marshal(vault)
	if err != nil {
		color.Red(err.Error())
		return
	}
	files.WriteIntoFile(data, "data.json")
	vault.UpdatedAt = time.Now()
}
