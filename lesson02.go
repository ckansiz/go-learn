package main

import (
	"fmt"
)

func main() {

	//for döngüsü

	// 1 ile 100 arasındaki sayıların 2 ile bölünenlerin toplamı
	total := 0
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			total += i
		}
	}

	fmt.Printf("2 ile bölünebilen sayıların toplamı : %d\n", total)

	//while döngüsü
	// 3 ile bölünebilen sayıların toplamı

	i := 0
	total = 0
	for i < 100 {
		if i%3 == 0 {
			total += i
		}
		i++
	}

	fmt.Printf("3 ile bölünebilen sayıların toplamı : %d\n", total)

	//sonsuz döngü kurmak istersek bu şekilde
	//for {

	//}

	var t1, t2, t3 int
	for i := 0; i < 100; i++ {
		if i%2 == 0 || i%3 == 0 {
			if i%2 == 0 {
				t2++
			}
			if i%3 == 0 {
				t3++
			}
		} else {
			t1++
		}

	}

	fmt.Printf("2 ile bölünebilen sayıların adedi : %d\t3 ile bölünebilen sayıların adedi : %d\tDigerleri : %d\n", t2, t3, t1)

	//switch case
	sinavNotu := 46
	switch {
	case sinavNotu >= 0 && sinavNotu < 45:
		fmt.Println("Üzgünüm Kaldın")
	case sinavNotu >= 45 && sinavNotu < 50:
		fmt.Println("Kanaat Kullanabilirim Galiba")
	case sinavNotu >= 50 && sinavNotu < 75:
		fmt.Println("Fena bir not değil")
	case sinavNotu >= 75 && sinavNotu <= 100:
		fmt.Println("Aferin..")
	default:
		fmt.Println("Böyle bir not YoKK")

	}

}
