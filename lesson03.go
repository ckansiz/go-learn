/*
	// Lesson_03
	Array, Slice, Map kullanımları
	Array'ler de boyut sabittir
	Slice'lar dinamik boyutludur. Kapasiteleri belirlenebilir.
	Map tipleri key:value veri modeline uygun şekilde kullanılır
	range fonskiyonu ile bir dizi,kesit veya harita'yı for döngüsü ile kullanabiliriz
	_ ile istersek bir fonksiyondan dönen değeri kullanmayacağımızı ifade edebiliriz
*/

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {

	// 81 il olarak şehirler dizisini set ettik
	var sehirler [81]string
	sehirler[34] = "istanbul"
	sehirler[6] = "Ankara"
	sehirler[16] = "Bursa"
	sehirler[35] = "İzmir"
	sehirler[61] = "Trabzon"

	// dizi uzunluğu
	fmt.Println(len(sehirler))

	// indis bazlı sehirler listesi
	for i := 0; i < len(sehirler); i++ {
		if sehirler[i] != "" {
			fmt.Printf("%s\t", sehirler[i])
		}
	}

	fmt.Println("")
	//foreach şehirler listesi (i indisi verir.)
	for i, sehir := range sehirler {
		if sehir != "" {
			fmt.Printf("%d %s\t", i, sehir)
		}
	}

	// 5 elemanlı float dizi
	adaylar := [5]float32{
		35,
		55,
		90,
		10,
		88,
	}

	var toplam float32
	var elemanSayisi = len(adaylar)
	fmt.Println(reflect.TypeOf(elemanSayisi)) // eleman_sayisi değişkeninin ismini yazdırdık
	for i := 0; i < elemanSayisi; i++ {
		toplam += adaylar[i]
	}

	ortalama := toplam / float32(elemanSayisi)
	fmt.Printf("Sınıfın ortalaması : %f\n", ortalama)

	var basarililar int = 0
	for _, puan := range adaylar {
		if puan >= 50 {
			basarililar++
		}
	}

	fmt.Printf("%d adet başarılı öğrenci var\n\n", basarililar)

	// 2 boyutlu dizi örneği

	matris := [3][2]int{{2, 3}, {6, 1}, {-9, 8}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\t", matris[i][j])
		}
		fmt.Println()
	}

	// örnek bir slice tanımııı
	oyuncu_adlari := []string{"mayk", "miki", "lora", "clara", "zorro", "dam", "edriyın", "raki", "barbarossa"}
	for _, value := range oyuncu_adlari {
		fmt.Println(value)
	}

	fmt.Println("Alt küme")

	alt_kume := oyuncu_adlari[6:]

	for _, value := range alt_kume {
		fmt.Println(value)
	}

	fmt.Println("eleman ekliyoruzzz")
	alt_kume = append(alt_kume, "hera")
	alt_kume = append(alt_kume, "sizar")

	for _, value := range alt_kume {
		fmt.Println(value)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var iller []string
	iller = make([]string, 5, 10) //5 eleman içeren 10a kadar genişleyebilen kesit. Eleman sayısı ve başlangıç kapasitesini belirttik
	iller[0] = "istanbul"
	iller[1] = "izmir"
	iller[2] = "ankara"
	iller[3] = "bursa"
	iller[4] = "gaziantep"
	iller = append(iller, "trabzon")

	for i := 0; i < len(iller); i++ {
		fmt.Printf("%d:%s\n", i, iller[i])
	}
	fmt.Println()
	for i, il := range iller {
		fmt.Printf("%d:%s\n", i, il)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	// map tanımlarıı

	sozluk := make(map[string]string)
	sozluk["white"] = "beyaz"
	sozluk["black"] = "siyah"
	sozluk["red"] = "kirmizi"
	sozluk["blue"] = "mavi"

	for key, value := range sozluk {
		fmt.Printf("[%s:%s]\n", key, value)
	}

	jsonContent, _ := json.MarshalIndent(sozluk, "", "")
	fmt.Println(string(jsonContent))

	//Bu sefer key içerikleri string value içerikleri int olan
	//bir map değişkeni tanımlandı
	envanter := map[string]int{
		"Laptop":       34,
		"Desktop":      5,
		"Tablet":       12,
		"Cep Telefonu": 34,
	}

	var toplam_cihaz int = 0
	for _, value := range envanter {
		toplam_cihaz += value
	}
	fmt.Printf("Envanterde %d cihaz var\n", toplam_cihaz)

}
