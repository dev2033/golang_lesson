package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// первый аргументы функции это то что функция будет принимать
// вторые скобки это то что функция будет возвращать

func GetFloatsSlice(filename string) ([]float64, error) {
	/*Нахождение среднего значение чисел из файла с числами*/
	var numbers []float64

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()

	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		if err != nil {
			return numbers, scanner.Err()
		}
	}
	return numbers, nil
}
