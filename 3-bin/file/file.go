package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	if !strings.HasSuffix(fileName, ".json") {
		return nil, errors.New("Неверный формат файла!")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись в файл успешно завершена!")
	file.Close()
}
