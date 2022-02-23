package main

import "fmt"

func main() {

	// Bir önceki derste de bahsettiğimiz gibi string tipi aslında bir byte dizisi barındırmaktadır.
	// Bir diğer öğrendiğimiz şey ise rune tipi aslında int32 değeri ile aynıdır ancak rune tipindeki bir değişkeni int32 tipindeki bir değere eşitleyemezsiniz.
	// Bunun için tip dönüşümü yapmanız gerekiyor.

	// rune karakteri metinsel ifade içerisindeki unicode karakterleri temsil etmektedir.
	// Go'da diğer dillerde olan char veri tipi bulunmamaktadır bunun yerine rune tipini kullanabilirsiniz.
	// Rune tipindeki değerler tek tırnak içerisine yazılmaktadır.
	// Tek tırnak içerisine sadece bir karakter yazabilirsiniz.
	r := '€'

	// aşağıdaki kodun çıktısı => int32, 114
	fmt.Printf("%T, %v\n", r, r)

	// rune tipindeki bir değeri string olarak yazmak isterseniz formatlarken %c ifadesini kullanmanız gerekmektedir.
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// bir string içerisindeki karakterleri for range ile almak istediğinizde her bir karakter rune tipinde gelecektir.
	// Bu karakterleri ekrana string olarak yazdırmak için %c ifadesini kullanabilirsiniz.
	myString := "Hello Patik@"
	for _, v := range myString {
		fmt.Printf("%c", v)
	}

}
