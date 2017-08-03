package main

import (
	"fmt"
)

func main() {
	var point int = 41

	fmt.Println(&point) //& ile değişkenin adresini alabiliriz.
	fmt.Printf("%d sayisinin bellek adresi %x dir\n", point, &point)

	pntr := &point
	fmt.Println(pntr, " adresindeki sayi ", *pntr)

	newLocation := pntr
	*newLocation = 99

	fmt.Println("Bellek adresleri pntr=>", pntr, " newLocation=>", newLocation)
	fmt.Println("Atama sonrası değerler\npntr=", *pntr, " newLocation=", *newLocation)

	luckyNumber := 7
	calculate(&luckyNumber)
	fmt.Println("Lucky Number : ", luckyNumber)

	stk := Stock{125, 25}
	fmt.Printf("high : %f low : %f", stk.high, stk.low)
	increaseStockByFifty(&stk)
	fmt.Printf("Stock min %f max %f\n", stk.low, stk.high)
}

func calculate(number *int) {
	//fonksiyonlara parametre olarak pointer verebiliriz.
	//Böylece fonksiyon içerisinde değişken kopyası oluşturmak yerine
	//referans adresi taşıdığımızdan daha optimize ve bellek dostu kod üretmiş olabiliriz
	fmt.Println("Fonksiyona gelen adres ", number, "\nDeğişken değeri ", *number)
	*number += 100

	somePoints := []float32{3.4, 2.1, 1.98, -4}
	doSomething(somePoints)
	for _, p := range somePoints {
		fmt.Println(p)
	}
}

// points bir slice ve referans türü
func doSomething(points []float32) {
	for i, point := range points {
		points[i] = point + 1
	}
}

type Stock struct {
	high float64
	low  float64
}

func increaseStockByFifty(stock *Stock) {
	stock.high += 50
	stock.low += 50
}
