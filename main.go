package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// inputName()
	// inputAge()
	// randomNumberGenerator()
	// arrayMethod()
	// readDataFromFile()

}

func readDataFromFile() {
	/*Чтение данных из файла*/
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func arrayMethod() {
	/*Массивы. Цикл*/
	notes := [7]string{
		"do",
		"re",
		"mi",
	}

	for _, value := range notes {
		fmt.Println(value)
	}

	// for item := 0; item < len(notes); item++ {
	// 	fmt.Println(notes[item])
	// }
}

func randomNumberGenerator() {
	/*
		Генерирует случаное число, преобразование строки
		в целочисленное число, циклы. Игра - угадай число.
	*/

	// создаем переменную, у которой значение в секундах
	seconds := time.Now().Unix()
	// инициализируем методом Seed() seconds
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я выбираю число от 1 до 100")
	fmt.Println("Число выбираю!")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("У вас осталось ", 10-guesses, " попыток")

		fmt.Print("Напишите ваше число: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		// метод Atoi() преобразует строку в целочисленное число
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Ваше значение больше чем загаданное")
		} else if guess < target {
			fmt.Println("Ваше значение меньше чем загаданное")
		} else {
			success = true
			fmt.Println("\nПоздравляю! Вы выиграли! \nЭто число - ", target)
			break
		}

	}

	if !success {
		fmt.Println("Вы проиграли! Ваши попытки закончились! Это число было: ", target)
	}
}

func inputAge() {
	/*Преобразование строки в число*/
	fmt.Print("Введите ваш возраст: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	// обрезает \n
	input = strings.TrimSpace(input)

	// преобразование ошибки нужно делать в 2х случаях
	if err != nil {
		log.Fatal(err)
	}

	// преобразуем значение
	old, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	if old >= 18 {
		fmt.Println("Ого! Какой ты взрослый!!!")
	} else {
		fmt.Println("Вам здесь не место")
	}

}

func inputName() {
	/*Прием данных от пользователя*/
	fmt.Print("Введите имя: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ваше имя - " + input)
}
