package main

import (
	"fmt"
	"strings"
)

func main() {
	colors := []string{
		"red", "blue", "green", "orange", "black", "yellow",
		"gray", "brown", "silver", "pink", "gold", "dark red",
	}

	// fonksiyonu içeride tanımlar isek;
	result := Select(colors, func(word string) bool {
		return len(word) >= 6
	})

	for _, word := range result {
		fmt.Printf("%s,", word)
	}

	fmt.Println()

	// önce fonksiyonu tanımlayıp sonra kullanır isek;
	g := func(w string) bool {
		return strings.HasPrefix(w, "g")
	}

	result = Select(colors, g)
	for _, word := range result {
		fmt.Printf("%s,", word)
	}

}

type predicate func(w string) bool

func Select(words []string, f predicate) []string {
	findings := []string{}
	for _, word := range words {
		if f(word) {
			findings = append(findings, word)
		}
	}
	return findings
}
