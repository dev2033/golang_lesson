/*
 Объявление переменных и вывод их
 Тут выводиться два типа перменных(числа и строки)
 с помощью Sprintf от модуля fmt. Чтобы посмотреть типы
 переменных можно прописать %T %T %T вместо %s %d %s

 Чтобы сложить один тип переменных с другим, нужно
 представить одну из них в тот тип с которой
 производим какие либо действия
*/

package main

import "fmt"

func main() {
	// textHello := "Hello"
	// textName := "John"

	i := 10

	// fmt.Println(fmt.Sprintf("%s %d %s", textHello, i, textName))

	var h float32 // еще один способ задания переменных
	h = 12.3

	next := float32(i) + h

	fmt.Println(next)

}
