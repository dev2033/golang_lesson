package main

import "fmt"

func main() {
	/* Управляющие конструкции if|else */
	a, b, c := 10, 20, 30

	if a > b {
		fmt.Println("a")
	} else if b > c {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

	var text = "Admin"
	switch text {
	case "admin":
		fmt.Println("Error!")
	case "Admin":
		fmt.Println("Yeah - Admin!")
		// выведет следующий кейс
		fallthrough
	case "Some admin":
		fmt.Println("Some case")
	default:
		fmt.Println("Some login")
	}

	switch {
	case text == "Admin" || text == "admin":
		fmt.Println("ok")
	default:
		fmt.Println("not ok")
	}

	// Циклы

	// Данный цикл выводит числа от 1 до 99, которые
	// делятся на 10 без остатка
	for x := 1; x < 100; x++ {
		if x%10 == 0 {
			fmt.Println(x)

		}
	}
}
