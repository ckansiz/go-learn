package main

import (
	"fmt"
	"time"
)

func main() {
	go loop1(0, 25)
	go loop2()
	go loop1(-25, -5)

	var enter string
	fmt.Println("Cikmak icin bir tusa basin")
	fmt.Scanln(&enter)
}

func loop1(min, max int) {
	for i := min; i < max; i++ {
		fmt.Printf("%d", i)
		time.Sleep(time.Microsecond * 500)
	}
}

func loop2() {
	var letters = "abcçdefgğhıijklmnoöprsştuüvyz"
	for _, c := range letters {
		fmt.Printf("%s", string(c))
		time.Sleep(time.Microsecond * 500)
	}
}
