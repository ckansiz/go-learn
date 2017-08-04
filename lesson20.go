package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	writeFileStats("file.png")
	fmt.Println()
	writeFileStats("NoFile.jpg")
	fmt.Println()
	writeToFile("golang.txt", "go inanılmaz keyifli bir dilmiş. Sanırım C kökenlilerin kolayca öğrenebileceği bir dil.")
	readFile("golang.txt")

	readFileWithIO("golang.txt")
	players := []Player{
		Player{1, "baltazar", 80},
		Player{2, "orvel", 23},
		Player{3, "nadya", 48},
		Player{4, "obi van", 91},
		Player{5, "şumi", 77},
	}

	f, err := os.Create("players.dat")
	if err == nil {
		var content string = "Players\n"
		for _, player := range players {
			content += player.ToString()
		}
		f.WriteString(content)
	} else {
		fmt.Printf("Dosya oluşturulurken hata meydana geldi\n%s\n", err.Error())
	}

	f, err = os.Create("players.json")
	if err == nil {
		jsonOutput, _ := json.Marshal(players)
		f.WriteString(string(jsonOutput))
	} else {
		fmt.Printf("Dosya oluşturulurken hata meydana geldi\n%s\n", err.Error())
	}

	count := dir("C:\\")
	fmt.Println(count, "dosya bulundu")
}

func (p Player) ToString() string {
	return strconv.Itoa(p.Id) + "|" + p.Title + "|" + strconv.Itoa(p.Level) + "\n"
}

type Player struct {
	Id    int
	Title string
	Level int
}

func readFileWithIO(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err == nil {
		fmt.Println(string(content))
	} else {
		fmt.Printf("Dosya okunmaya çalışılırken hata meydana geldi\n%s\n", err.Error())
	}
}

func readFile(path string) {
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		fileInfo, _ := f.Stat()
		fileBytes := make([]byte, fileInfo.Size())
		read, _ := f.Read(fileBytes)
		fmt.Printf("\n%s (%d btye okundu)\n\t%s\n", path, read, fileBytes)
	} else {
		fmt.Printf("Dosyayı açamadım dostum\n\t%s\n", err.Error())
	}
}

func writeToFile(path, text string) {
	f, err := os.Create(path)
	if err == nil {
		defer f.Close()
		r, _ := f.WriteString(text)
		fmt.Println("İşlem sonucu", r)
	} else {
		fmt.Printf("Dosyaya yazmada hata\n%s\n", err.Error())
	}
}

func writeFileStats(path string) {
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		fInfo, _ := f.Stat()
		fmt.Println(fInfo.Size(), "bytes")
		fmt.Println("File Name is", fInfo.Name())
		fmt.Println("File Modes are", fInfo.Mode().String())
		fmt.Println(fInfo.ModTime(), "Last Changed")
	} else {
		fmt.Printf("Ben bu dosyayı açamadım\n%s\n", err.Error())
	}
}

//bir klasör içeriğini gösterelim
func dir(path string) int {
	directory, err := os.Open(path)
	var fileCount int = 0
	if err == nil {
		defer directory.Close()
		files, _ := directory.Readdir(-1)
		for _, f := range files {
			if !f.IsDir() {
				fmt.Printf("%s\t\t%d\t%s\n", f.Name(), f.Size(), f.ModTime())
				fileCount++
			}
		}
	} else {
		fmt.Printf("Klasör okumada sorun oluştu\n\t%s\n", err.Error())
	}
	return fileCount
}
