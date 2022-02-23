package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Bildiğiniz gibi string ifadeler birer byte array olarak saklanmaktadır.
	// Bir sayıyı string bir ifadeye çevirmek için string() fonksiyonunu kullanabilirsiniz.

	s := string(100) // d karakterinin ASCII gösterimi 100 olduğu için sonuç d olacak.
	fmt.Println(s)

	// Ancak bu işlemi strconv paketi ile yaptığınızda parametre olarak verilen int değerini olduğu gibi string'e çevirecek
	myVar := strconv.Itoa(100)
	fmt.Println(myVar)
}
