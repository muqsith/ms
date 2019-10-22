package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dirName := "."
	if len(os.Args) > 1 {
		dirName = os.Args[1]
	}
	files, err := ioutil.ReadDir(dirName)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println("Name: ", file.Name(), "\nModified: ", file.ModTime(), "\n----------")
	}
}
