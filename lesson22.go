package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Sunucu dinlemede...")
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/players", getPlayers)      // /players şeklinde gelen talepleri alan fonksiyon getPlayers
	log.Fatal(http.ListenAndServe(":4501", nil)) //http://localhost:4501 üzerinden dinlemede
}

func mainPage(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("[%s]\t%s:%s\n", request.Method, time.Now(), request.URL)
	response.Header().Set("Content-Type", "text/html")
	f, _ := ioutil.ReadFile("mainPage.html")
	io.WriteString(response, string(f))
}

func getPlayers(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("[%s]\t%s:%s\n", request.Method, time.Now(), request.URL) //log basıyoruz
	players := []Player{                                                 //Örnek bir slice oluşturuyoruz
		Player{1, "Con Doo", 100},
		Player{2, "Wolfy", 89},
		Player{3, "Red Sonia", 71},
		Player{4, "Dam Klause", 54},
		Player{5, "Turko", 99},
	}
	json.NewEncoder(response).Encode(players) //json encoder ile players'ı responseWriter üzerinden yolluyoruz
}

type Player struct {
	Id       int     "json:`ID`"
	Nickname string  "json:`Nickname`"
	Level    float32 "json:`Level`"
}
