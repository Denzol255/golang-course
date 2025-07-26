package account

import (
	"app/password/encrypter"
	"encoding/json"
	"slices"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write(data []byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDB struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func (vault *VaultWithDB) DeleteAccountByUrl(url string) bool {
	accountIdx := slices.IndexFunc(vault.Accounts, func(account Account) bool {
		return account.Url == url
	})
	if accountIdx != -1 {
		vault.Accounts = append(vault.Accounts[:accountIdx], vault.Accounts[accountIdx+1:]...)
		vault.save()
	}
	return accountIdx != -1
}

func (vault *VaultWithDB) FindAccounts(checkerFunc func(account Account) bool) []Account {
	result := make([]Account, 0)
	for _, value := range vault.Accounts {
		if checkerFunc(value) {
			result = append(result, value)
		}
	}

	return result
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  make([]Account, 0),
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	var vault Vault
	data := enc.Decrypt(file)
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		color.Red(err.Error())
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  make([]Account, 0),
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDB) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.save()
}

func (vault *VaultWithDB) save() {
	data, err := json.Marshal(vault.Vault)
	if err != nil {
		color.Red(err.Error())
		return
	}
	encryptedData := vault.enc.Encrypt(data)
	vault.db.Write(encryptedData)
	vault.UpdatedAt = time.Now()
}
