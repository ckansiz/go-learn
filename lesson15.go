package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.Second * 10 // programın timeout süresi ayarlanıyor
	fmt.Println("Eş zamanlı iş için zaman aşımı süresi", timeout)

	// cn1 adında channel oluşturuluyor
	cn1 := make(chan string, 1)
	go someJob(cn1)

	// channel sonuçları bekleniyor
	// bir önceki örnekteki for ile işlemlerin sonucu bekleniyordu. Onun yerine burada time.After ile timeout süresi kadar bekleme yapılıyor
	select {
	case result := <-cn1:
		fmt.Println(result)
	case <-time.After(timeout):
		fmt.Println("İş istenilen sürede tamamlanamadı!")
	}

	timeout = 10
	fmt.Println("Yeni zaman aşımı(Timeout) süresi", timeout)
	jobCn1 := make(chan string, 1)
	counterCn1 := make(chan bool, 1)

	go someJob(jobCn1)
	go counter(counterCn1, timeout)
	select {
	case jobMsg := <-jobCn1:
		fmt.Println(jobMsg)
	case counterMsg := <-counterCn1:
		fmt.Println("Zaman aşımı oluşma durumu", counterMsg)
	}
}

func someJob(msg chan string) {
	fmt.Println("\tEş zamanlı işler yapılıyor")
	time.Sleep(time.Second * 5) // işin yapılma süresi 5 sn
	msg <- "İş tamamlandı"
}

func counter(result chan bool, duration time.Duration) {
	fmt.Println("Sayaç", duration, "kadar bekleyecek")
	time.Sleep(time.Second * duration)
	result <- true
}
