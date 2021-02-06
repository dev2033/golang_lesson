package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// первый аргументы функции это то что функция будет принимать
// вторые скобки это то что функция будет возвращать
func GetFloats(filename string) ([3]float64, error) {
	/*Нахождение среднего значение чисел из файла с числами*/

	var numbers [3]float64

	file, err := os.Open(filename)

	if err != nil {
		return numbers, err
	}

	i := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
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
