package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func countLetters(url string, frequency *map[string]int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	sentence := string(body)
	words := strings.Split(sentence, " ")
	for _, wordIncludeSpace := range words {
		word := strings.ReplaceAll(wordIncludeSpace, " ", "")
		word = strings.ReplaceAll(word, "\n", "")
		if word == "" {
			continue
		}
		if _, ok := (*frequency)[word]; !ok {
			(*frequency)[word] = 1
		} else {
			(*frequency)[word]++
		}
	}

	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1010; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, &frequency)
	}

	time.Sleep(10 * time.Second)

	for k, v := range frequency {
		fmt.Printf("%s-%d ", k, v)
	}
	fmt.Println()
}
