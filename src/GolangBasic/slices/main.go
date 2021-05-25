package main

import "fmt"

func main() {
	var sliceOne []int
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	sliceOne = append(sliceOne, 10)
	sliceOne = append(sliceOne, 20)
	sliceOne = append(sliceOne, 30)
	sliceOne = append(sliceOne, 40)
	sliceOne = append(sliceOne, 50)
	sliceOne = append(sliceOne, 60)
	sliceOne = append(sliceOne, 70)

	// fmt.Println(len(sliceOne))
	// fmt.Println(cap(sliceOne))
	fmt.Println(sliceOne[1:6])

	slice := make([]int, 5)
	fmt.Println(len(slice))
}
