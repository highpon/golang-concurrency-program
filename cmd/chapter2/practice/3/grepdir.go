package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getAllFiles(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func catFile(fileName, searchString string) error {
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

	allFiles := getAllFiles(os.Args[2])

	searchString := os.Args[1]
	fmt.Println("allFiles", allFiles, os.Args[2])
	for _, fileName := range allFiles {
		go catFile(fileName, searchString)
	}

	time.Sleep(5 * time.Second)
}
