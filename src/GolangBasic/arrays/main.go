package main

import "fmt"

func main() {
	/*Массивы*/
	var array [5]int
	array[0] = 10
	array[1] = 513
	array[2] = 32
	array[3] = 31
	array[4] = 32
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for i, v := range array {
		fmt.Printf("Индекс - %v, Значение - %v\n", i, v)
	}
}
