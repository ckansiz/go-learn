package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	paragraph := `Line 1 : If there was a problem walking to the file or directory named by path, 
	Line 2 : the incoming error will describe the problem and the function can decide 
	Line 3 : how to handle that error. If an error is returned, processing stops. 
	Line 4 : The sole exception is when the function returns the special value SkipDir.
	Line 5:end of code !` // " yerine ` kullanarak birden fazla satırdan oluşan metinler belirtebiliriz

	fmt.Println(paragraph)
	containsSample("an")
	containsSample("ist")
	countSample(paragraph)
	fieldsSample(paragraph, 7)
	fieldsFuncSample(paragraph)

	sampleText := "bir bilmecem var çocuklar! Haydi sor sor, çay'da kahvaltı da yenir :) Acaba nedir nedir?"
	hasPrefixSample(sampleText, "bir")
	hasPrefixSample(sampleText, "peki")

	indexNo := strings.IndexAny(sampleText, ":,?!;.") // indexAny verilen karakterlerden ilk bulduğunu "0" indexten sayarak verir.
	fmt.Println(indexNo)

	values := []string{"data source", "tcp", "connection", "log", "function"}
	newValue := strings.Join(values, "|") // values isimli string dizideki elemanların arasına | işareti koyarak birleştirir
	fmt.Println(newValue)

	// Join benzeri eğlenceli fonksiyonlardan birisi de Repeat
	// Bir ifadenin n sayıda tekrarı için kullanılıyor.
	fmt.Printf("Hey Ney%s\n", strings.Repeat("Na", 3))

	rot13 := func(c rune) rune { //rune tipinin anlamlandığı bir yer. int32 gibi olan rune aslında karakterin sayısal karşılığını veriyor.
		return c + 13
	}
	encryptedText := strings.Map(rot13, "ordu sağ kanattan sabah şafakla harekete geçecek")
	fmt.Println(encryptedText)
	decryptedText := strings.Map(func(c rune) rune {
		return c - 13
	}, encryptedText)
	fmt.Println(decryptedText)

	//Split fonksiyonu ile bir metni belli bir karaktere göre ayırmamız mümkündür
	product := "Pro Go Lang|Book|35,50|550|ISBN:345676"
	columns := strings.Split(product, "|")
	for _, column := range columns {
		fmt.Println(column)
	}

	// Title, ToLower, ToUpper ile harf çevirimleri
	motto := "bEniM hala UMUDUM vaAArr"
	fmt.Println(motto)
	fmt.Println(strings.Title(motto))   //sadece başharfleri büyüğe çevirdi
	fmt.Println(strings.ToUpper(motto)) // tüm harfleri büyüye çevirdi
	fmt.Println(strings.ToLower(motto)) // tüm harfleri küçüğe çevirdi
}

func hasPrefixSample(text string, searching string) {
	if strings.HasPrefix(text, searching) {
		fmt.Printf("\n'%s' in başında '%s' VAR!\n", text, searching)
	} else {
		fmt.Printf("\n'%s' in başında '%s' YOK!\n", text, searching)
	}
}

func fieldsFuncSample(paragraph string) {
	words := strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsNumber(c)
	})
	fmt.Printf("%q", words)
}

func fieldsSample(paragraph string, charCount int) {
	fmt.Println("Harf Sayısı -", charCount, "- dan fazla olanlar")
	words := strings.Fields(paragraph)
	for _, word := range words {
		if len(word) >= charCount {
			fmt.Println(word)
		}
	}
}

func countSample(part string) {
	spaceCount := strings.Count(part, " ")
	lineCount := strings.Count(part, "\n")
	fmt.Printf("Metinde %d boşluk ve %d satır bulundu. Buna göre toplam kelime sayısı %d\n", spaceCount, lineCount, spaceCount+lineCount-1)
}

func containsSample(part string) {

	words := []string{
		"istanbul",
		"mavi",
		"yeşil",
		"antalya",
		"pist",
		"ankara",
		"izmir",
		"antartika",
		"arktika",
		"anason",
		"mandalina",
		"banana",
	}

	fmt.Println("'", part, "' geçen kelimeler")
	for _, word := range words {
		if strings.Contains(word, part) {
			fmt.Println(word)
		}
	}

}
