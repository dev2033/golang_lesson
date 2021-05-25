package main

import "fmt"

func main() {
	var simpleMap = map[int]string{}
	simpleMap[198] = "some word"
	fmt.Println(simpleMap)

	var anotherMap = make(map[string]int)
	anotherMap["one"] = 11
	anotherMap["two"] = 22
	anotherMap["three"] = 33
	fmt.Println(anotherMap)

	delete(anotherMap, "three")
	fmt.Println(anotherMap)

	slice := []int{1, 2, 3}
	for _, v := range slice {
		fmt.Println(v)
	}

	// Проверяет есть ли такой ключ в map,
	// если есть, то выводит значение по этому ключу
	if value, ok := anotherMap["one"]; ok {
		fmt.Println(value)
	}
}
