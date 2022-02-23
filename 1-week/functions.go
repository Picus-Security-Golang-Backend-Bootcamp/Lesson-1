// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	message, age, err := printMessage("Patika", 17465489, "Turkey")

// 	if err != nil {
// 		//loglayabilirsiniz.
// 		// kullanıcıya hata mesajı dönebilirsiniz
// 		fmt.Println(err)
// 	}

// 	fmt.Println(message)
// 	fmt.Println(age)

// }

// func printMessage(name string, age int, nation string) (string, int, error) {
// 	message := fmt.Sprintf("Hello, %s", name)
// 	//veri tabanı kodları
// 	var err error
// 	if age < 18 {
// 		err = errors.New("Hata oluştu")
// 		return "", age, err
// 	}
// 	return message, age, nil
// }
