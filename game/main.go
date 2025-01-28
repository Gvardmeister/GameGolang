package main

import (
	"fmt"
	"strings"
)

func (s *location) formatLocation() {
	fmt.Printf("Ваша текущая локация: %s\n", s.currentLocation)

	var listLocation []string
	for _, value := range s.goLocation {
		listLocation = append(listLocation, value)
	}
	fmt.Println("Доступные локации:", strings.Join(listLocation, " "))
}

// var commandsGlobalMap = map[int]string{
// 1:  "идти коридор",
// 2:  "идти комната",
// 3:  "идти кухня",
// 4:  "идти улица",
// 5:  "осмотреться",
// 6:  "надеть рюкзак",
// 7:  "взять ключи",
// 8:  "взять конспекты",
// 9:  "взять телефон",
// 10: "применить ключи дверь",
// 11: "применить телефон шкаф",
// 12: "применить ключи шкаф",
// 13: "завтракать",
// }

// // значение в мапе - это кол-во предметов
// var itemMap = map[string]int{
// 	"рюкзак":    1,
// 	"ключи":     1,
// 	"конспекты": 1,
// 	"телефон":   1,
// }

var commands = []string{
	"идти коридор",
	"идти комната",
	"идти кухня",
	"идти улица",
	"осмотреться",
	"надеть рюкзак",
	"взять ключи",
	"взять конспекты",
	"взять телефон",
	"применить ключи дверь",
	"применить телефон шкаф",
	"завтракать",
}

type person struct {
	commandPerson string // возможно нужно изменить
	// inventory     string
}

type location struct {
	// itemRoom  int если реализую механизм генерации предмета
	currentLocation string
	goLocation      []string
}

func main() {
	fmt.Println("В игре Вам доступны локации - [улица, коридор, комната, кухня]")
	fmt.Println()

	var locat string

	fmt.Println("Выберите текущую локацию: ")
	fmt.Scan(&locat)

	switch locat {
	case "улица", "коридор", "комната", "кухня":
		fmt.Println()
		fmt.Printf("Вы выбрали локацию - %s\n", locat)
		fmt.Println()
	default:
		fmt.Println()
		fmt.Println("Такой локации не существует")
		fmt.Println()
		// сгенерировать ошибку и выключить игру
		break
	}

	switch locat {
	case "улица":
		var street = location{
			currentLocation: "улица",
			goLocation:      []string{"коридор"}, // почему нельзя указать без типа данных - обязательно указываем тип
		}
		street.formatLocation()
	case "коридор":
		var corridor = location{
			currentLocation: "коридор",
			goLocation:      []string{"комната", "кухня"},
		}
		corridor.formatLocation()
	case "комната":
		var room = location{
			currentLocation: "комната",
			goLocation:      []string{"коридор"},
		}
		room.formatLocation()
	case "кухня":
		var kitchen = location{
			currentLocation: "кухня",
			goLocation:      []string{"коридор"},
		}
		kitchen.formatLocation()
	}

	fmt.Println()
	fmt.Println("Нажмите {1} - для просмотра доступных команд\nИли {2} - чтобы выйти из меню")

	var check int
	fmt.Scan(&check)

	fmt.Println()
	fmt.Printf("Вы выбрали - %d", check)
	fmt.Println()

	if check == 1 {
		fmt.Println()
		fmt.Println("Список доступных команд:")
		for _, value := range commands {
			fmt.Println(value)
		}
	} else if check == 2 {
		fmt.Println("Выход из меню команд")
	} else {
		fmt.Println("Введена неправильная команда")
	}
}
