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

	"./datafile"
)

func main() {
	// InputName()
	// InputAge()
	// RandomNumberGenerator()
	// ArrayMethod()
	// ReadDataFromFile()
	// AverageValueOfNumbers()
	// Slice()
	// AverageValueOfNumbersSlice()
	// Map()

}

func Struct() {
	/*Тип данных - Структуры*/
	var subscriber struct {
		name   string
		rate   float64
		active bool
	}
	subscriber.name = "Alex"
	subscriber.rate = 11.00
	subscriber.active = true
}

func Map() {
	/*
		Подсчитывает сколько в файле data3 было
		одинаковых слов(колличество написанных ЯП)
	*/
	lines, err := datafile.GetStrings("data3.txt")

	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}

func AverageValueOfNumbersSlice() {
	/*Подсчет среднего значения чисел при помощи сегментов*/
	numbers, err := datafile.GetFloatsSlice("data2.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	count := float64(len(numbers))
	fmt.Println(sum / count)
}

func Slice() {
	/*Сегменты*/
	// arrayLiters := [5]string{
	// 	"a",
	// 	"b",
	// 	"c",
	// 	"d",
	// 	"e",
	// }
	// arrayLiters[0] = "Q"
	// slice := arrayLiters[0:3] // [:2] - альтернатива
	// slice2 := arrayLiters[2:]
	// fmt.Println(slice)
	// fmt.Println(slice2)

	// ---------------------------------------------------

	slice := []string{
		"q",
		"w",
	}
	// Добавление
	slice = append(slice, "e")
	fmt.Println(slice)

}

func AverageValueOfNumbers() {
	/*
		Вызывает функцию из пакета datafile для дальнейшей
		обработке, после эту функцию вызываем в main
	*/
	numbers, err := datafile.GetFloats("data2.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	count := float64(len(numbers))
	fmt.Println(sum / count)
}

func ReadDataFromFile() {
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

func ArrayMethod() {
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

func RandomNumberGenerator() {
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

func InputAge() {
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

func InputName() {
	/*Прием данных от пользователя*/
	fmt.Print("Введите имя: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ваше имя - " + input)
}
