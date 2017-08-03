package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	m1 := make(chan string)
	m2 := make(chan string)
	m3 := make(chan string)

	go jobA(m1, "A..C")
	go jobB(m1, "C..M")
	go jobC(m1, "M..Z")

	for i := 0; i < 3; i++ {
		select {
		case messageA := <-m1:
			fmt.Println(messageA)
		case messageB := <-m2:
			fmt.Println(messageB)
		case messageC := <-m3:
			fmt.Println(messageC)
		}
	}

	fmt.Println(time.Now())

}

func jobA(msg chan string, sets string) {
	fmt.Println(sets, "için işlemler yapılıyor")
	time.Sleep(time.Second * 2)
	msg <- "A görevi tamamlandı"
}

func jobB(msg chan string, sets string) {
	fmt.Println(sets, "için işlemler yapılıyor")
	time.Sleep(time.Second * 9)
	msg <- "B görevi tamamlandı"
}

func jobC(msg chan string, sets string) {
	fmt.Println(sets, "için işlemler yapılıyor")
	time.Sleep(time.Second * 6)
	msg <- "C görevi tamamlandı"
}
