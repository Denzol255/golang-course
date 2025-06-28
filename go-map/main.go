package main

import "fmt"

type bookMarkMap = map[string]string

func main() {
	fmt.Println("Ваши ссылки: ")
	links := make(bookMarkMap)

Menu:
	for {
		fmt.Println("")
		fmt.Println("- 1. Посмотреть ссылки")
		fmt.Println("- 2. Добавить ссылку")
		fmt.Println("- 3. Удалить ссылку")
		fmt.Println("- 4. Выход")

		var choice int

		fmt.Print("Ваш выбор: ")
		fmt.Scanf("%d", &choice)
		fmt.Println("")

		switch choice {
		case 1:
			showLinks(links)
		case 2:
			key := getLinkKey()
			value := getLinkValue()
			links[key] = value
			fmt.Println("Ссылка успешно добавлена")
		case 3:
			key := getLinkKey()
			delete(links, key)
			fmt.Println("Ссылка успешно удалена")
		case 4:
			break Menu
		default:
			fmt.Println("Некорректный выбор")
		}
	}
}

func showLinks(links bookMarkMap) {
	if len(links) == 0 {
		fmt.Println("Список ссылок пуст")
		return
	}
	fmt.Println("Список ссылок:")
	for key, value := range links {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func getLinkKey() string {
	var key string
	fmt.Println("Укажите название ссылки: ")
	fmt.Scanln(&key)
	return key
}

func getLinkValue() string {
	var value string
	fmt.Println("Укажите ссылку: ")
	fmt.Scanln(&value)
	return value
}
