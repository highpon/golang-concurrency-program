package main

import (
	"fmt"
	"os"
	"time"
)

func catFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		return err
	}
	fmt.Println(string(data[:count]))

	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		go catFile(arg)
	}

	time.Sleep(5 * time.Second)
}
