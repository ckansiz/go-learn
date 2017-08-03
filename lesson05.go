package main

import (
	"fmt"
)

func main() {

	fmt.Println("main baslangici")
	sonuc := dosyaSifrele("bilgiler.dat")
	fmt.Println(sonuc)
	fmt.Println("main bitisi")
	fmt.Println(dosyaAyristir("urunler.json"))
}

func dosyaSifrele(dosya string) bool {
	defer dosyayiKapat(dosya)
	defer bellegiTemizle()
	fmt.Println("Şifeleme operasyonu yapılıyor")
	return true
}

func dosyayiKapat(dosya string) {
	fmt.Println("Kalan veriler dosyaya yazilip kapatiliyor")
}

func bellegiTemizle() {
	fmt.Println("Bellek Temizleniyor")
}

func dosyaAyristir(dosya string) string {
	defer func() {
		fmt.Printf("%s için gerekli kapatma operasyonları yapılacaktır\n", dosya)
	}()
	fmt.Println("Dosya açılıyor...")
	fmt.Println("Ayrıştırma işlemi yapılıyor")
	return "operasyon başarılı"
}
