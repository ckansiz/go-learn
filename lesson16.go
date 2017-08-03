package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 4)
	fmt.Println("timer nesnesi tanımlandı. Kod Akışına devvam ediyor")
	fmt.Println(time.Now())
	now := <-timer.C
	fmt.Println("Timer ile belirtilen süre doldu")
	fmt.Println(now)

	timer = time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("İkinci timer süresi geçti")
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("Timer Durduruldu")
	}

	tickTime := time.NewTicker(time.Second * 2)
	go func() {
		fmt.Println("İş yapıyorum")
		for t := range tickTime.C {
			fmt.Println(t)
		}
	}()

	var enter string
	fmt.Println("Çıkmak için Enter tuşuna basınız")
	fmt.Scanln(&enter)
	tickTime.Stop()
}
