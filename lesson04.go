package main

import "fmt"

func main() {
	fmt.Println("Fonksiyonları tanıyalım")
	toplam := Topla(4, 5)
	fmt.Println("4+5 = ", toplam)
	a, b, c, d := DortIslem(3, 2)
	fmt.Printf("%f\t%f\t%f\t%f\n", a, b, c, d)
	fmt.Println(CokluToplam(1, 2, 3, 4, 5))
	sayiSlice := []int{3, 7, 1, 9, 10}
	fmt.Println(SliceToplam(sayiSlice))

	sozluk := map[string]string{
		"black": "kara",
		"white": "beyaz",
		"gold":  "altın",
	}
	map_yazdir(sozluk)
	fmt.Println(Faktoryel(10))

	var puan int
	fmt.Println("Aldığınız puanı girer misiniz?")
	fmt.Scanln(&puan)

	puan_kontrol_func := func(d int) bool {
		if d < 50 {
			return false
		}
		return true
	}

	fmt.Println(puan_kontrol_func(puan))
}

func Topla(a, b int) int {
	return a + b
}

func DortIslem(a, b float32) (toplam, carpim, bolum, fark float32) {

	toplam = a + b
	carpim = a * b
	bolum = a / b
	fark = a - b

	return toplam, carpim, bolum, fark
}

func CokluToplam(sayilar ...int) int {
	toplam := 0
	for _, i := range sayilar {
		toplam += i
	}

	return toplam
}

func SliceToplam(sayilar []int) (toplam int) {
	for _, i := range sayilar {
		toplam += i
	}

	return toplam
}

func map_yazdir(sozluk map[string]string) {
	for k, v := range sozluk {
		fmt.Printf("[%s:%s]\n", k, v)
	}
}

func Faktoryel(sayi int) (faktoryel int) {

	if sayi == 0 || sayi == 0 {
		faktoryel = 1
	} else {
		faktoryel = sayi * Faktoryel(sayi-1)
	}

	return faktoryel
}
