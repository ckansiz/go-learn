package main

import (
	"fmt"
	"reflect"
)

func main() {

	ad := "jan\tclaud\n"
	var soyad string = "van\t"

	z := ad + soyad

	fmt.Println(z)

	var isim, parola, email bool
	isim = true
	parola = false
	email = true

	fmt.Println(isim && parola && email)
	fmt.Println(isim || parola || email)
	fmt.Println("İstanbul" == "Ankara")

	const pi float32 = 3.1415
	fmt.Println(pi)

	birDeger := false
	fmt.Println(reflect.TypeOf(pi), reflect.TypeOf(birDeger))

}
