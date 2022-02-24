package main

import "fmt"

func main() {

	// Pazar günü işlediğimiz derste fonksiyonları öğrendik ve birkaç örnek ile öğrendiklerimizi uygulamaya döktük.
	// ancak genel olarak bir fonksiyonu bir değişkene atıp kullandık ya da herhangi bir değişkene ihtiyacımız olmayacak şekilde kullandık (örn : sadece ekrana mesaj yazan fonksiyon)

	// Fonksiyonlar eğer tek bir dönüş değerine sahipse hiçbir değişkene atmadan başka bir metodun/fonksiyonun parametresi olarak kullanabilirsiniz.

	fmt.Println(getMessage())

	//---------------------------------------------------------------------

	// "defer" kullanımı
	deferUsage()

	//---------------------------------------------------------------------

	// Fonksiyonlarınızı da bir değişkene atayarak kullanabilirsiniz
	x := getMessage
	x()

	// Eğer değişkene atadığınız fonksiyonunuz bir parametre alıyorsa aşağıdaki gibi kullanabilirsiniz
	//x(parametre)

	//---------------------------------------------------------------------

	// Anonim fonksiyonlar
	// Anonim fonksiyonlar çoğu kaynakta function literal, lambda function ya da closure olarak bahsedilir. En basit tanımı ile tanımladığımız tüm isimsiz fonksiyonlar anonim fonksiyondur. Go projelerinin hemen hemen %90'ınında bu tarz fonksiyonları görebilirsiniz.

	anonFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println()
	fmt.Printf("Anonim fonksiyon sonucu : %d\n", anonFunc(5, 5))

	// anonim fonksiyonların bir diğer kullanım şekli ise tanımlandığı anda çağırmak. Eğer JS biliyorsanız bu IIFE (Immediately Invoked Function Expression) ile aynı mantıkta. => https://www.javascripttutorial.net/javascript-immediately-invoked-function-expression-iife/

	secondAnonFunc := func(a, b int) int {
		return a + b
	}(5, 5) // parametre değerlerimizi burada veriyoruz ve sonucu direkt olarak değişkene atıyoruz

	fmt.Printf("İkinci anonim fonksiyon sonucu : %d\n", secondAnonFunc)

	// Günün sonunda aklınızda tutmanız gereken kalıp ise :
	// Eğer anonim bir fonksiyon tanımlıyorsanız => func() {...}
	// Eğer tanımladığınız anonim fonksiyonu anında çalıştırmak istiyorsanız => func() {...}()

	// Ekstra bilgi : Go makalelerinin çoğunda "functions are first-class citizens" ifadesini görürsünüz bunun anlamı şu; fonksiyonlar başka fonksiyonlara parametre olarak kullanılabilir, dönüş değeri olarak kullanılabilir ve bir değişkene atanabilirler.

	// ------------------------------------------------------------------------
	// Fonksiyonları parametre olarak kullanmak
	// Yukarıda da bahsettiğim gibi bir fonksiyonu başka bir fonksiyonun parametresi olarak kullanabilirsiniz ve bir fonksiyondan başka bir fonksiyon dönebilirsiniz.

	// Bunun için öncelikle anonim bir fonksiyon tanımladık ve bir değişkene atadık
	sumFunc := func(a, b int) int {
		return a + b
	}

	// Ardından tanımladığımız fonksiyonu parametre olarak başka bir fonksiyona geçtik
	funcParameter(sumFunc)

	// Yukarıda yaptığımız işlemin bir diğer yazım şekli aşağıdaki gibidir

	funcParameter(func(a, b int) int {
		return a + b
	})
}

func funcParameter(f func(int, int) int) {
	fmt.Println()
	fmt.Printf("Parametre olarak alınan fonksiyonun sonucu : %d\n", f(5, 5))
}

func deferUsage() {

	// Eğer bir fonksiyon içerisinde fonksiyon bitmeden hemen önce yapacağınız bir işlem varsa bunun için Go'da defer anahtar kelimesini kullanabilirsiniz.
	// defer en basit haliyle fonksiyonunuz bittiğinde çalışır.
	// defer hakkında bilmeniz gereken bir diğer şey ise defer ne olursa olsun çalışır. İster fonksiyonunuzda fatal bir hata olsun ister if-else bloklarının birinde hata fırlatın ya da geriye değer dönün defer her zaman için fonksiyon bitmeden hemen önce çalışır

	// Fonksiyonunuzda birden fazla defer kullanımı olabilir. Bu durumda en son yazdığınızdan ilk yazdığınıza doğru çalışır.
	result := 5 * 10
	defer fmt.Printf("Defer ile ekrana yazdırdığımız ilk sonuç : %d\n", result)
	defer fmt.Printf("Defer ile ekrana yazdırdığımız ikinci sonuç : %d\n", result)
	defer fmt.Printf("Defer ile ekrana yazdırdığımız üçüncü sonuç : %d\n", result)

	// defer olmayan diğer kodlar sıralı bir şekilde çalışmaya devam eder
	fmt.Println("Defer'den önce çalışacak olan ilk method")
	fmt.Println("Defer'den önce çalışacak olan ikinci method")
	fmt.Println("Defer'den önce çalışacak olan üçüncü method")

	fmt.Println()
}

func getMessage() string {
	return "Hello world"
}
