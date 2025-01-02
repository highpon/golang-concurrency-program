package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func grepFile(fileName, searchString string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		return err
	}

	if strings.Contains(string(data[:count]), searchString) {
		fmt.Printf("%s contains %s.\n", fileName, searchString)
		return nil
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		panic("number of args is invalid")
	}

	searchString := os.Args[1]
	for _, arg := range os.Args[2:] {
		go grepFile(arg, searchString)
	}

	time.Sleep(5 * time.Second)
}
