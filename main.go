package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/gonutz/auto"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	text := strings.Join(os.Args[1:], " ")
	if len(os.Args) == 1 {
		text, _ = auto.ClipboardText()
	}

	words := strings.Split(text, " ")
	for i := range words {
		word := []byte(words[i])
		var start, end int
		for i := range word {
			if unicode.IsLetter(rune(word[i])) {
				start = i
				break
			}
		}
		for i := len(word) - 1; i >= 0; i-- {
			if unicode.IsLetter(rune(word[i])) {
				end = i
				break
			}
		}
		if end-start >= 3 {
			shuffle(word[start+1 : end])
		}
		words[i] = string(word)
	}

	auto.SetClipboardText(strings.Join(words, " "))
}

func shuffle(b []byte) {
	n := len(b)
	for i := 0; i < n-1; i++ {
		j := i + rand.Intn(n-i)
		if i != j {
			b[i], b[j] = b[j], b[i]
		}
	}
}
