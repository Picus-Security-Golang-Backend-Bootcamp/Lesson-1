package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Go'da tip dönüşümlerini iki şekilde yapabilirsiniz.
	// İlk yöntem string(), int(), float64() gibi fonksiyonları kullanarak

	// string() fonksiyonu için dikkat etmeniz gereken nokta ise 100 vb. karakterleri verdiğinizde sonucunda 100 değerini string tipinde almazsınız.
	// Bunun sebebi string() fonksiyonu parametre olarak verilen int değerinin ASCII tablosundaki temsil ettiği karaktere çevirir
	s := string("my string")
	fmt.Printf("string() fonsksiyonu : %s\n", s)

	// 100 değeri tam sayı olmasına rağmen float64() fonksiyonunu kullanarak bu değerin tipini int'ten float64'e çevirdik
	myFloat := float64(100)
	// float tipindeki değerleri formatlama esnasında %f ile gösterebilirsiniz.
	// Eğer noktadan sonraki kısmı kısaltmak isterseniz %.2f ya da %.3f olarak formatlayabilirsiniz.
	fmt.Printf("float64() fonksiyonu : %f", myFloat)

	// Bir diğer yöntem ise string tipingdeki değerleri int,float64 vb. tiplere dönüştürmek için kullanabileceğimiz strconv paketi.

	// Burada string() fonksiyonundan farklı olarak "d" karakterini değil metinsel ifade olarak 100 değerini alıyoruz.

	// strconv paketi tip dönüşümlerinde 2 dönüş değerine sahiptir. İlki verdiğiniz değerin dönüştürmek istediğiniz tipteki hali, ikincisi de error tipinde değer

	// Ancak Itoa metodu (Integer to ASCII) string tipine çevirdiğinden dolayı error tipinde bir dönüş değerine sahip değildir. Verdiğiniz tüm değerleri string'e çevirebilir.
	myString := strconv.Itoa(100)

	fmt.Printf("strconv.Itoa metodu : %s", myString)

	// Itoa hariç geri kalan fonksiyonlar ise belirli bir geçerliliğe sahip olması koşulundan dolayı iki adet dönüş değeri içerir.
	// ParseFloat metodunun ikinci parametresi ise dönüştürmek istediğiniz tipin float32 mi yoksa float64 mü olacağını belirlemeniz için kullanılır

	// Burada myFloat hala float64 olsa da float32 tipine değeri değiştirilmeden dönüştürülebilir. Bu özellikle noktadan sonraki kısımda yer alan sayılar önemli olduğunda dikkat etmeniz gereken bir konu
	myFloat, err := strconv.ParseFloat("sdfs", 32)

	if err != nil {
		fmt.Println("Hata")
	}

	fmt.Printf("strconv.ParseFloat metodu : %f", myFloat)

}
