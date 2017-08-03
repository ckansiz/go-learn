package main

import (
	"fmt"
	"time"
)

func main() {

	transitInfo := make(chan string)
	go func() {
		fmt.Println("Burası mapping işlemlerini yapar")
		time.Sleep(time.Millisecond * 1000)
		transitInfo <- "mapping sonucu 3 donör bulundu"
	}()

	incomingInfo := <-transitInfo
	fmt.Printf("Map isimli goroutine'den '%s' bilgisi döndü\n", incomingInfo)
	close(transitInfo) // kanalı kapatıyoruz

	// 2. örnek
	message := make(chan string)
	go Listener(message)
	input := <-message
	fmt.Printf("Listener dedi ki [%s]\n", input)
	close(message)

	//3. örnek
	parts := make(chan string, 5)
	parts <- "ayakkabılar"
	parts <- "ceketler"
	parts <- "pantalonlar"
	parts <- "bluzlar"
	parts <- "çoraplar"

	for i := 0; i < 5; i++ {
		go func(p chan string) {
			value := <-p
			fmt.Printf("\t[%s] parçalar işlenecek\n", value)
		}(parts)
	}

	close(parts)

	// 4. örnek
	// kanal oluştur
	soundChannel := make(chan string, 1)
	// kanal'dan HOLA! şeklinde bir mesaj al
	microphone(soundChannel, "HOLA!")
	// kanalı işlemek üzere tek yönlü soundBox fonksiyonuna gönder
	soundBox(soundChannel, 10, 1500)

	// 5. örnek
	statusChannel := make(chan bool)
	go worker(statusChannel)
	<-statusChannel

	var enter string
	fmt.Printf("\nÇıkmak için Enter\n")
	fmt.Scanln(&enter)
}

func worker(completed chan bool) {
	fmt.Println("Bir takım işler yapılıyor...")
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("İşlemler tamamlandı")
	completed <- true
}

// kanaldan veri alma şeklinde çalışan bir channel
func microphone(sound chan<- string, message string) {
	sound <- message
}

// kanala veri göndermek için çalışan bir channel örneği
func soundBox(sound <-chan string, volumeLevel int, duration int) {
	fmt.Printf("Sound is %s.\nLevel = %d\nDuration = %d\n", <-sound, volumeLevel, duration)
}

func Listener(msg chan string) {
	msg <- "pong"
	go func(chn chan string) {
		chn <- "ping"
	}(msg)

	output := <-msg
	fmt.Printf("\t iç fonksiyon dedi ki [%s]\n", output)
}
